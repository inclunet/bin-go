package bingo

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	LoadRounds(context.Context) ([]*Round, error)
	SaveRound(context.Context, *Round) error
	SaveRounds(context.Context, ...*Round) error
}

type PostgresStore struct {
	pool *pgxpool.Pool
}

func NewPostgresStore(pool *pgxpool.Pool) Store {
	if pool == nil {
		return nil
	}

	return &PostgresStore{pool: pool}
}

func (s *PostgresStore) LoadRounds(ctx context.Context) ([]*Round, error) {
	rows, err := s.pool.Query(ctx, `
		SELECT id, display_number, type, next_round_id
		FROM bingo_rounds
		ORDER BY display_number
	`)
	if err != nil {
		return nil, fmt.Errorf("load bingo rounds: %w", err)
	}
	defer rows.Close()

	rounds := make([]*Round, 0)
	roundByID := make(map[string]*Round)

	for rows.Next() {
		var (
			id          uuid.UUID
			nextRoundID pgtype.UUID
			round       = &Round{}
		)

		if err := rows.Scan(&id, &round.Round, &round.Type, &nextRoundID); err != nil {
			return nil, fmt.Errorf("scan bingo round: %w", err)
		}

		round.ID = id.String()
		if nextRoundID.Valid {
			nextID := uuid.UUID(nextRoundID.Bytes)
			round.NextRoundID = nextID.String()
		}
		round.Cards = []Card{}
		rounds = append(rounds, round)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate bingo rounds: %w", err)
	}

	for i := range rounds {
		roundByID[rounds[i].ID] = rounds[i]
	}

	cardRows, err := s.pool.Query(ctx, `
		SELECT id, round_id, display_number, state
		FROM bingo_cards
		ORDER BY round_id, display_number
	`)
	if err != nil {
		return nil, fmt.Errorf("load bingo cards: %w", err)
	}
	defer cardRows.Close()

	for cardRows.Next() {
		var (
			id            uuid.UUID
			roundID       uuid.UUID
			displayNumber int
			state         []byte
			card          Card
		)

		if err := cardRows.Scan(&id, &roundID, &displayNumber, &state); err != nil {
			return nil, fmt.Errorf("scan bingo card: %w", err)
		}
		if err := json.Unmarshal(state, &card); err != nil {
			return nil, fmt.Errorf("decode bingo card %s: %w", id, err)
		}

		card.ID = id.String()
		card.RoundID = roundID.String()
		card.Card = displayNumber

		round := roundByID[card.RoundID]
		if round == nil {
			return nil, fmt.Errorf("bingo card %s references unknown round %s", card.ID, card.RoundID)
		}
		round.Cards = append(round.Cards, card)
	}
	if err := cardRows.Err(); err != nil {
		return nil, fmt.Errorf("iterate bingo cards: %w", err)
	}

	drawnByRound := make(map[string]map[int]bool)
	drawRows, err := s.pool.Query(ctx, `SELECT round_id, number FROM bingo_draws`)
	if err != nil {
		return nil, fmt.Errorf("load bingo draws: %w", err)
	}
	defer drawRows.Close()

	for drawRows.Next() {
		var (
			roundID uuid.UUID
			number  int
		)
		if err := drawRows.Scan(&roundID, &number); err != nil {
			return nil, fmt.Errorf("scan bingo draw: %w", err)
		}
		id := roundID.String()
		if drawnByRound[id] == nil {
			drawnByRound[id] = make(map[int]bool)
		}
		drawnByRound[id][number] = true
	}
	if err := drawRows.Err(); err != nil {
		return nil, fmt.Errorf("iterate bingo draws: %w", err)
	}

	for i := range rounds {
		sort.Slice(rounds[i].Cards, func(a, b int) bool {
			return rounds[i].Cards[a].Card < rounds[i].Cards[b].Card
		})
		restoreDrawnNumbers(rounds[i], drawnByRound[rounds[i].ID])
		rounds[i].RestoreRuntime()
	}

	return rounds, nil
}

func restoreDrawnNumbers(round *Round, drawn map[int]bool) {
	if len(round.Cards) == 0 {
		return
	}

	main := &round.Cards[0]
	main.Checked = 0
	for line := range main.Numbers {
		for column := range main.Numbers[line] {
			number := &main.Numbers[line][column]
			number.Checked = drawn[number.Number]
			if number.Checked {
				main.Checked++
			}
		}
	}
	main.Finished = main.IsFinished()
}

func (s *PostgresStore) SaveRound(ctx context.Context, round *Round) error {
	return s.SaveRounds(ctx, round)
}

func (s *PostgresStore) SaveRounds(ctx context.Context, rounds ...*Round) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin saving bingo rounds: %w", err)
	}
	defer tx.Rollback(context.Background()) //nolint:errcheck

	for _, round := range rounds {
		if err := saveRound(ctx, tx, round); err != nil {
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit bingo rounds: %w", err)
	}

	return nil
}

func saveRound(ctx context.Context, tx pgx.Tx, round *Round) error {
	roundID, err := uuid.Parse(round.ID)
	if err != nil {
		return fmt.Errorf("parse bingo round id %q: %w", round.ID, err)
	}

	var nextRoundID interface{}
	if round.NextRoundID != "" {
		parsed, err := uuid.Parse(round.NextRoundID)
		if err != nil {
			return fmt.Errorf("parse next bingo round id: %w", err)
		}
		nextRoundID = parsed
	}

	if _, err := tx.Exec(ctx, `
		INSERT INTO bingo_rounds (id, display_number, type, next_round_id)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE SET
			display_number = EXCLUDED.display_number,
			type = EXCLUDED.type,
			next_round_id = EXCLUDED.next_round_id,
			updated_at = NOW()
	`, roundID, round.Round, round.Type, nextRoundID); err != nil {
		return fmt.Errorf("save bingo round %s: %w", round.ID, err)
	}

	for i := range round.Cards {
		card := &round.Cards[i]
		cardID, err := uuid.Parse(card.ID)
		if err != nil {
			return fmt.Errorf("parse bingo card id %q: %w", card.ID, err)
		}
		state, err := json.Marshal(card)
		if err != nil {
			return fmt.Errorf("encode bingo card %s: %w", card.ID, err)
		}

		if _, err := tx.Exec(ctx, `
			INSERT INTO bingo_cards (id, round_id, display_number, state)
			VALUES ($1, $2, $3, $4)
			ON CONFLICT (id) DO UPDATE SET
				display_number = EXCLUDED.display_number,
				state = EXCLUDED.state,
				updated_at = NOW()
		`, cardID, roundID, card.Card, state); err != nil {
			return fmt.Errorf("save bingo card %s: %w", card.ID, err)
		}
	}

	drawnNumbers := make([]int32, 0)
	if len(round.Cards) > 0 {
		for _, line := range round.Cards[0].Numbers {
			for _, number := range line {
				if !number.Checked || number.Number <= 0 {
					continue
				}
				drawnNumbers = append(drawnNumbers, int32(number.Number))
			}
		}
	}

	if _, err := tx.Exec(ctx, `
		DELETE FROM bingo_draws
		WHERE round_id = $1
			AND NOT (number = ANY($2::integer[]))
	`, roundID, drawnNumbers); err != nil {
		return fmt.Errorf("remove stale bingo draws for round %s: %w", round.ID, err)
	}

	for _, number := range drawnNumbers {
		if _, err := tx.Exec(ctx, `
			INSERT INTO bingo_draws (round_id, number)
			VALUES ($1, $2)
			ON CONFLICT DO NOTHING
		`, roundID, number); err != nil {
			return fmt.Errorf("save bingo draw for round %s: %w", round.ID, err)
		}
	}

	return nil
}
