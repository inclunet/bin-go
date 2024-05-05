package braille

import "github.com/inclunet/bin-go/pkg/translator"

type Challenge struct {
	Challenge string
	Repply    string
	Word      string
}

func (c *Challenge) Check(repply string) bool {
	return c.Word == translator.ToggleString(repply)
}

func NewChallenge(word string) Challenge {
	challenge := "word"

	if GetRandomNumber(0, 1) == 1 {
		word = translator.ToggleString(word)
		challenge = "braille"
	}

	return Challenge{
		Repply:    "",
		Challenge: challenge,
		Word:      word,
	}
}
