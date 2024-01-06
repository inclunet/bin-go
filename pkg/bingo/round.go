package bingo

import (
	"fmt"
)

type Round struct {
	Cards  []Card
	Round  int
	rounds *Rounds
	Type   int
}

func (r *Round) CheckNumber(card, number int) *Card {
	currentCard, err := r.GetCard(card)

	if err != nil {
		return nil
	}

	mainCard, err := r.GetCard(0)

	if err != nil {
		return currentCard
	}

	if !mainCard.IsChecked(number) && card > 0 {
		return currentCard
	}

	if card == 0 {
		r.CheckNumberForAll(number)
	}

	currentCard.CheckNumber(number)

	return currentCard
}

func (r *Round) CheckNumberForAll(number int) *Round {
	for i, card := range r.Cards {
		if card.Autoplay && i > 0 {
			r.Cards[i].CheckNumber(number)
		}
	}

	return r
}

func (r Round) Draw() *Card {
	mainCard, err := r.GetCard(0)

	if err != nil {
		return nil
	}

	mainCard.Draw()

	return mainCard
}

func (r *Round) GetCard(card int) (*Card, error) {
	if card < 0 || card >= len(r.Cards) || len(r.Cards) == 0 {
		return nil, fmt.Errorf("Card %d not found", card)
	}

	return &r.Cards[card], nil
}

func (r *Round) SetNextRound(round int) *Round {
	if round == 0 || round == r.Round {
		return r
	}

	oldRound, err := r.rounds.GetRound(round)

	if err != nil {
		return r
	}

	for card := range oldRound.Cards {
		oldRound.Cards[card].SetNextRound(r.Round)
	}

	return r
}

func (r *Round) ToggleAutoplay(card int) *Card {
	currentCard, err := r.GetCard(card)

	if err != nil {
		return nil
	}

	currentCard.ToggleAutoplay().CheckDrawedNumbers()

	return currentCard
}

func (r *Round) UncheckNumber(number int) *Round {
	for i := range r.Cards {
		r.Cards[i].uncheckNumber(number)
	}

	return r
}

func NewRound(rounds *Rounds, roundType int) (round Round) {
	round.rounds = rounds
	round.Round = len(rounds.Rounds) + 1
	round.Type = roundType
	NewCard(&round)
	round.rounds.Rounds = append(round.rounds.Rounds, round)
	return round
}
