package bingo

type Rounds struct {
	Total  int
	Rounds []Round
}

func (r *Rounds) AddCard(round int) *Card {
	return r.Rounds[round].AddCard()
}

func (r *Rounds) AddRound(roundType int) *Card {
	id := len(r.Rounds) + 1
	round := NewRound(id, roundType)
	card := round.AddCard()
	r.Rounds = append(r.Rounds, round)
	return card
}

func (r *Rounds) CheckNumber(round, card, number int) *Card {
	if r.Rounds[round].Cards[0].IsChecked(number) || card == 0 {
		return r.Rounds[round].Cards[card].CheckNumber(number)
	}

	return &r.Rounds[round].Cards[card]
}

func (r *Rounds) GetCard(round, card int) *Card {
	return r.GetRound(round).GetCard(card)
}

func (r *Rounds) GetRound(round int) *Round {
	return &r.Rounds[round]
}

func (r *Rounds) ToggleNumber(round, card, number int) *Card {
	if r.GetCard(round, card).IsChecked(number) && card == 0 {
		return r.UncheckNumber(round, card, number)
	}

	return r.CheckNumber(round, card, number)
}

func (r *Rounds) UncheckNumber(round, card, number int) *Card {
	if r.GetCard(round, 0).IsChecked(number) && card == 0 {
		return r.GetRound(round).UncheckNumber(number).GetCard(card)
	}

	return r.GetCard(round, card)
}
