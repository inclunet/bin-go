package bingo

type Round struct {
	Cards []Card
	Round int
	Type  int
}

func (r *Round) AddCard() *Card {
	card := Card{
		Card:  len(r.Cards) + 1,
		Round: r.Round,
		Type:  r.Type,
	}

	card.DrawCard()
	r.Cards = append(r.Cards, card)
	return &card
}

func (r *Round) CheckNumber(card, number int) *Card {
	if card == 0 {
		return r.CheckNumberForAll(number).GetCard(card)
	}

	return r.GetCard(card).CheckNumber(number)
}

func (r *Round) CheckNumberForAll(number int) *Round {
	for i, card := range r.Cards {
		r.Cards[i].LastNumber = number

		if card.AutoPlay {
			r.Cards[i].CheckNumber(number)
		}
	}

	return r
}

func (r Round) Draw() *Card {
	return r.CheckNumberForAll(r.GetCard(0).Draw().LastNumber).GetCard(0)
}

func (r *Round) GetCard(card int) *Card {
	return &r.Cards[card]
}

func (r *Round) ToggleAutoplay(card int) *Card {
	return r.GetCard(card).ToggleAutoplay().CheckDrawedNumbers(*r.GetCard(0))
}

func (r *Round) UncheckNumber(number int) *Round {
	for i := range r.Cards {
		r.Cards[i].uncheckNumber(number)
	}

	return r
}

func NewRound(id int, bingoType int) Round {
	return Round{Round: id, Type: bingoType}
}
