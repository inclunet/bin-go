package braille

import (
	"encoding/json"
	"os"

	"github.com/inclunet/bin-go/pkg/utils"
)

var classes []Class

type Class struct {
	Description         string
	RequiredPunctuation int
	Rounds              int
	AllowedChallenges   []string
	Words               []string
}

func (c *Class) DrawAllowedChallenge() string {
	t := len(c.AllowedChallenges)

	if t == 0 {
		return ""
	}

	r := utils.GetRandomNumber(0, t-1)

	return c.AllowedChallenges[r]
}

func (c *Class) DrawWord() string {
	t := len(c.Words)

	if t == 0 {
		return ""
	}

	r := utils.GetRandomNumber(0, t-1)

	return c.Words[r]
}

func GetClass(i int) *Class {
	c := len(classes)

	if c == 0 {
		return nil
	}

	if i < 0 {
		i = 0
	}

	if i >= c {
		i = c - 1
	}

	return &classes[i]
}

func LoadClass(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&classes)

	return err
}
