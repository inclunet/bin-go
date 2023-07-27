package bingo

type Round struct {
	Cards []Card
	Round int
	Type  int
}

func (r *Round) AddCard() Card {
	card := Card{
		Card:  len(r.Cards) + 1,
		Round: r.Round,
		Type:  r.Type,
	}

	card.Draw()
	r.Cards = append(r.Cards, card)
	return card
}

func (r *Round) GetBingoCard(card int) Card {
	return r.Cards[card-1]
}

func (r *Round) uncheckNumber(number int) *Round {
	for i := range r.Cards {
		r.Cards[i].uncheckNumber(number)
	}

	return r
}

func NewRound(id int, bingoType int) Round {
	return Round{Round: id, Type: bingoType}
}
