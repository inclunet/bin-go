package tictac

import "fmt"

type Round struct {
	Card    *Card
	Players []Player
}

func (r Round) GetPlayer(id int) (*Player, error) {
	if id < 0 {
		return nil, fmt.Errorf("invalid player id: %d", id)
	}

	if id >= len(r.Players) {
		return nil, fmt.Errorf("invalid player id: %d", id)
	}

	return &r.Players[id], nil
}

func NewRound(round, t int) *Round {
	return &Round{
		Card:    NewCard(round, t),
		Players: []Player{},
	}
}
