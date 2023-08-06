package bingo

type Rounds struct {
	Total  int
	Rounds []Round
}

func (r *Rounds) AddCard(round int) *Card {
	r.Total++
	return r.GetRound(round).AddCard()
}

func (r *Rounds) AddRound(roundType int) *Card {
	id := len(r.Rounds) + 1
	round := NewRound(id, roundType)
	card := round.AddCard()
	card.AutoPlay = true
	r.Rounds = append(r.Rounds, round)
	return card
}

func (r *Rounds) CheckNumber(round, card, number int) *Card {
	if r.GetCard(round, 0).IsChecked(number) || card == 0 {
		return r.GetCard(round, card).CheckNumber(number)
	}

	return r.GetCard(round, card)
}

func (r *Rounds) Draw(round int) *Card {
	return r.GetRound(round).Draw()
}

func (r *Rounds) GetCard(round, card int) *Card {
	return r.GetRound(round).GetCard(card)
}

func (r *Rounds) GetRound(round int) *Round {
	return &r.Rounds[round]
}

func (r *Rounds) ToggleAutoplay(round, card int) *Card {
	return r.GetRound(round).ToggleAutoplay(card)
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
