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

func (r *Round) AddCard() *Card {
	card := NewCard(*r)

	r.Cards = append(r.Cards, card)

	return &card
}

func (r *Round) CheckNumberForAll(number int) int {
	counter := 0

	for i, card := range r.Cards {
		if card.Autoplay && i > 0 {
			if r.Cards[i].CheckNumber(number) {
				counter++
			}
		}
	}

	return counter
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

func (r *Round) SetNextRound(nextRound int) bool {
	if nextRound == 0 || nextRound == r.Round {
		return false
	}

	for card := range r.Cards {
		if r.Cards[card].SetNextRound(nextRound) {
			return true
		}
	}

	return false
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
		if card.Autoplay && i > 0 {
			if card.IsChecked(number) {
				if r.Cards[i].UncheckNumber(number) {
					uncheckCounter++
				}
			} else {
				if r.Cards[i].CheckNumber(number) {
					checkCounter++
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

func NewRound(bingo Bingo, roundType int) (round Round) {
	round.Round = len(bingo.Rounds) + 1
	round.Type = roundType

	round.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	round.Cards = append(round.Cards, NewCard(round))

	return round
}
