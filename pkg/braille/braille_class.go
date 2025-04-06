package braille

import (
	"time"
)

type BrailleClass struct {
	Challenge          Challenge
	class              *Class
	CurrentClass       int
	CurrentPunctuation int
	CurrentRound       int
	Description        string
	Player             int
	PreviousClass      int
	NextClass          int
	StartedAt          time.Time
	TotalRounds        int
	TotalPunctuation   int
	UpdatedAt          time.Time
}

func (b *BrailleClass) Check(repply string) *BrailleClass {
	b.CurrentRound++

	if b.Challenge.Check(repply) {
		b.CurrentPunctuation++
		b.TotalPunctuation++
		b.GetChallenge()

		return b
	}

	b.TotalPunctuation--

	return b
}

func (b *BrailleClass) GetClass() *BrailleClass {
	b.CurrentPunctuation = 0
	b.CurrentRound = 1
	b.PreviousClass = 0

	if b.CurrentClass > 0 {
		b.PreviousClass = b.CurrentClass - 1
	}

	b.NextClass = b.CurrentClass

	if b.CurrentClass < len(classes)-1 {
		b.NextClass = b.CurrentClass + 1
	}

	if class := GetClass(b.CurrentClass); class != nil {
		b.class = class
		b.Description = b.class.Description
		b.TotalRounds = b.class.Rounds
	}

	b.GetChallenge()

	return b
}

func (b *BrailleClass) GetChallenge() *BrailleClass {
	if b.CurrentRound > b.class.Rounds {
		if b.CurrentPunctuation >= b.class.RequiredPunctuation {
			b.CurrentClass++
		}

		return b.GetClass()
	}

	b.Challenge = NewChallenge(b.class.DrawWord())

	b.UpdatedAt = time.Now()

	return b
}

func NewBrailleClass(player int) (b BrailleClass) {
	b.StartedAt = time.Now()
	b.Player = player
	b.GetClass()
	b.UpdatedAt = time.Now()

	return b
}
