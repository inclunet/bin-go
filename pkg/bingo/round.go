package bingo

import "log"

type Round struct {
	Cards []Card
	Round int
	Type  int
}

func (r *Round) AddBingoCard() Card {
	card := Card{
		Card:  len(r.Cards) + 1,
		Round: r.Round,
		Type:  r.Type,
	}
	card.Draw()
	log.Printf("Adding card %d into %d round\n", card.Card, card.Round)
	r.Cards = append(r.Cards, card)
	return card
}

func (r *Round) GetBingoCard(card int) Card {
	return r.Cards[card-1]
}

func NewRound(id int, bingoType int) Round {
	return Round{Round: id, Type: bingoType}
}
