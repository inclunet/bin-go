package bingo

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Round struct {
	Cards    []Card
	Round    int
	Type     int
	upgrader websocket.Upgrader
}

func (r *Round) AddCard() (*Card, error) {
	card := NewCard(r)

	r.Cards = append(r.Cards, card)

	return r.GetCard(card.Card - 1)
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

func (r *Round) SetCompletionsForAll(completions *Completions) (int, error) {
	counter := 0

	if len(r.Cards) == 0 {
		return counter, fmt.Errorf("no cards found")
	}

	for i := range r.Cards {
		if err := r.Cards[i].SetCompletions(completions); err != nil {
			return counter, err
		}

		counter++
	}

	return counter, nil
}

func (r *Round) SetNextRoundForAll(nextRound int) int {
	count := 0

	if nextRound < 0 {
		return count
	}

	for card := range r.Cards {
		if r.Cards[card].SetNextRound(nextRound) {
			count++
		}
	}

	return count
}

func (r *Round) SetRoundForAll(round *Round) bool {
	for card := range r.Cards {
		r.Cards[card].SetNextRound(round.Round)
	}

	return false
}

func (r *Round) ToggleAutoplay(card int) *Card {
	currentCard, err := r.GetCard(card)

	if err != nil {
		return nil
	}

	// currentCard.ToggleAutoplay().CheckDrawedNumbers()

	return currentCard
}

func (r *Round) ToggleNumberForAll(number int) (int, int) {
	checkCounter := 0
	uncheckCounter := 0

	for i, card := range r.Cards {
		if i > 0 {
			if card.Autoplay && !card.IsChecked(number) {
				if r.Cards[i].CheckNumber(number) {
					checkCounter++
				}
			} else {
				if r.Cards[i].UncheckNumber(number) {
					uncheckCounter++
				}
			}
		}
	}

	return checkCounter, uncheckCounter
}

func (r *Round) UncheckNumberForAll(number int) *Round {
	for i := range r.Cards {
		if i > 0 {
			r.Cards[i].UncheckNumber(number)
		}
	}

	return r
}

func NewRound(bingo *Bingo, roundType int) Round {
	round := Round{
		Round: len(bingo.Rounds) + 1,
		Type:  roundType,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		Cards: []Card{},
	}

	round.AddCard()

	return round
}
