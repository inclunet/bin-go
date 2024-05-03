package classes

import "github.com/inclunet/bin-go/pkg/translator"

type Challenge struct {
	Braille []translator.BrailleCell
	Word    string
}

func (c *Challenge) Check(challenge Challenge) bool {
	if translator.BrailleToString(c.Braille) != challenge.Word {
		return false
	}

	if translator.BrailleToString(challenge.Braille) != c.Word {
		return false
	}

	return true
}

func NewChallenge(word string) Challenge {
	return Challenge{
		Braille: translator.StringToBraille(word),
		Word:    word,
	}
}
