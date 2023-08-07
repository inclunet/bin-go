package bingo

type Round struct {
	Cards []Card
	Round int
	Type  int
}

func (r *Round) AddCard() *Card {
	card := NewCard(r.Round, len(r.Cards)+1, r.Type)
	r.Cards = append(r.Cards, card)

	if r.GetCard(0).Checked == 0 {
		r.GetCard(card.Card - 1).ToggleAutoplay()
	}

	return r.GetCard(card.Card - 1)
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

func (r *Round) GetCard(card int) (defaultCard *Card) {
	//if card < len(r.Cards) {
	return &r.Cards[card]
	//}

	//return defaultCard
}

func (r *Round) SetNextRound(round int) *Round {
	for card := range r.Cards {
		r.GetCard(card).SetNextRound(round)
	}

	return r
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

func NewRound(round, roundType int) (newRound Round) {
	newRound.Round = round
	newRound.Type = roundType

	newRound.AddCard()

	return newRound
}
