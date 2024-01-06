package bingo

import (
	"fmt"
	"log"
)

type Rounds struct {
	Total  int
	Rounds []Round
}

func (r *Rounds) AddCard(round int) Card {
	currentRound, err := r.GetRound(round)

	if err != nil {
		return Card{}
	}

	r.Total++

	return NewCard(currentRound)
}

func (r *Rounds) AddRound(oldRoundId, roundType int) (*Card, error) {
	log.Println("Adding round", oldRoundId, roundType)

	currentRound := NewRound(r, roundType)
	log.Println("Adding round", oldRoundId, roundType)
	mainCard, err := currentRound.GetCard(0)

	log.Println("Adding round", oldRoundId, roundType)

	if err != nil {
		return nil, err
	}

	oldRound, err := r.GetRound(oldRoundId)

	if err != nil {
		return mainCard, nil
	}

	if oldRoundId > 0 {
		oldRound.SetNextRound(currentRound)
	}

	return mainCard, nil
}

func (r *Rounds) CheckNumber(round, card, number int) *Card {
	currentRound, err := r.GetRound(round)

	if err != nil {
		return nil
	}

	return currentRound.CheckNumber(card, number)
}

func (r *Rounds) Draw(round int) *Card {
	currentRound, err := r.GetRound(round)

	if err != nil {
		return nil
	}

	return currentRound.Draw()
}

func (r *Rounds) GetCard(round, card int) *Card {
	currentRound, err := r.GetRound(round)

	if err != nil {
		return nil
	}

	currentCard, err := currentRound.GetCard(card)

	if err != nil {
		return nil
	}

	return currentCard
}

func (r *Rounds) GetRound(round int) (*Round, error) {
	if round < 0 || round >= len(r.Rounds) || len(r.Rounds) == 0 {
		return nil, fmt.Errorf("Round %d not found", round)
	}

	return &r.Rounds[round], nil
}

func (r *Rounds) ToggleAutoplay(round, card int) *Card {
	currentRound, err := r.GetRound(round)

	if err != nil {
		return nil
	}

	return currentRound.ToggleAutoplay(card)
}

func (r *Rounds) ToggleNumber(round, card, number int) *Card {
	if r.GetCard(round, card).IsChecked(number) && card == 0 {
		return r.UncheckNumber(round, card, number)
	}

	return r.CheckNumber(round, card, number)
}

func (r *Rounds) UncheckNumber(round, card, number int) *Card {
	currentRound, err := r.GetRound(round)

	if err != nil {
		return nil
	}

	mainCard, err := currentRound.GetCard(0)

	if err != nil {
		return nil
	}

	if mainCard.IsChecked(number) && card == 0 {
		currentRound.UncheckNumber(number)

		return mainCard
	}

	return r.GetCard(round, card)
}
