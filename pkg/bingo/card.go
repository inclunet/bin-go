package bingo

import (
	"sort"

	"github.com/gorilla/websocket"
	"github.com/inclunet/bin-go/pkg/utils"
)

type Card struct {
	main           *Card
	Completions    Completions
	Autoplay       bool
	Bingo          bool
	Round          int
	Card           int
	conn           *websocket.Conn
	Checked        int
	LastCompletion string
	LastNumber     int
	NextRound      int
	Type           int
	Numbers        [][5]Number
}

func (c *Card) CancelAlert() {
	if c.Bingo {
		c.Bingo = false
		c.UpdateCard()
	}
}

func (c *Card) CheckDrawedNumbers(main *Card) (counter int) {
	if !c.Autoplay {
		return counter
	}

	if main.Checked == 0 {
		return 0
	}

	for _, line := range main.Numbers {
		for _, number := range line {
			if number.Checked {
				if c.CheckNumber(number.Number) {
					counter++
				}
			}
		}
	}

	return counter
}

func (c *Card) CheckNumber(number int) bool {
	if c.main != nil {
		if !c.main.IsChecked(number) {
			return false
		}
	}

	for l, line := range c.Numbers {
		for col, column := range line {
			if column.Number == number && !column.Checked {
				c.LastNumber = number
				c.Numbers[l][col].Checked = true
				c.Checked++
				c.Bingo, c.LastCompletion = c.IsBingo()
				c.UpdateCard()
				return true
			}
		}
	}

	return false
}

func (c *Card) Draw() int {
	if c.Card > 1 {
		return 0
	}

	if c.Checked >= c.Type {
		return 0
	}

	newNumber := utils.GetRandomNumber(1, c.Type)

	if !c.CheckNumber(newNumber) {
		return c.Draw()
	}

	return newNumber
}

func (c *Card) DrawCard() {
	c.Numbers = c.GetEmptyBingoCard()

	for col := 1; col <= 5; col++ {
		for l, number := range c.drawColumn(col) {
			c.Numbers[l][col-1] = number
		}
	}

	if c.Card > 1 {
		c.Numbers[2][2].Checked = true
		c.Numbers[2][2].Number = 0
	}
}

func (c *Card) drawColumn(column int) []Number {
	var numbers []Number

	for l := 1; l <= c.GetCardLines(); l++ {
		if c.Card == 1 {
			lines := c.Type / 5
			numbers = append(numbers, Number{Checked: false, Column: column, Number: l + (lines*column - lines)})
		} else {
			numbers = append(numbers, c.drawNumber(numbers, column))
		}
	}

	if c.Card > 1 {
		numbers = c.ShortNumbers(numbers)
	}

	return numbers
}

func (c *Card) drawNumber(drawedNumbers []Number, column int) Number {
	columnEnd := column * c.Type / 5
	columnStart := columnEnd - (c.Type/5 - 1)
	drawNumber := Number{Checked: false, Column: column, Number: utils.GetRandomNumber(columnStart, columnEnd)}

	if IsDuplicatedNumber(drawedNumbers, drawNumber) && drawNumber.Number >= columnStart && drawNumber.Number <= columnEnd {
		return c.drawNumber(drawedNumbers, column)
	}

	return drawNumber
}

func (c *Card) GetCardLines() int {
	if c.Card == 1 {
		return c.Type / 5
	} else {
		return 5
	}
}

func (c *Card) GetEmptyBingoCard() [][5]Number {
	var bingoCard [][5]Number

	for i := 0; i < c.GetCardLines(); i++ {
		bingoCard = append(bingoCard, [5]Number{})
	}

	return bingoCard
}

func (c *Card) IsBingo() (bool, string) {
	if c.main == nil {
		return false, ""
	}

	if c.IsFull() {
		return true, "Full"
	}

	if c.IsHorizontal() {
		return true, "Horizontal"
	}

	if c.IsVertical() {
		return true, "Vertical"
	}

	if c.IsDiagonal() {
		return true, "Diagonal"
	}

	return false, ""
}

func (c *Card) IsDiagonal() bool {
	c.Completions.Diagonal.Update(&c.main.Completions.Diagonal)

	if c.main.Completions.Diagonal.GetRemaining() == 0 {
		return false
	}

	if c.Completions.Diagonal.GetRemaining() == 0 {
		return false
	}

	counter := 0

	for l, r, cl, cr := 0, 4, 0, 0; l < 5; l, r = l+1, r-1 {
		if c.Numbers[l][l].Checked {
			cl++

			if cl == 5 {
				counter++
			}
		}

		if c.Numbers[l][r].Checked {
			cr++

			if cr == 5 {
				counter++
			}
		}
	}

	if !c.Completions.Diagonal.Check(counter) {
		return false
	}

	if !c.main.Completions.Diagonal.Add() {
		return false
	}

	return true
}

func (c *Card) IsFull() bool {
	c.Completions.Full.Update(&c.main.Completions.Full)

	if c.main.Completions.Full.GetRemaining() == 0 {
		return false
	}

	if c.Completions.Full.GetRemaining() == 0 {
		return false
	}

	if c.Checked < 24 {
		return false
	}

	if !c.Completions.Full.Check(1) {
		return false
	}

	if !c.main.Completions.Full.Add() {
		return false
	}

	return true
}

func (c *Card) IsHorizontal() bool {
	c.Completions.Horizontal.Update(&c.main.Completions.Horizontal)

	if c.main.Completions.Horizontal.GetRemaining() == 0 {
		return false
	}

	if c.Completions.Horizontal.GetRemaining() == 0 {
		return false
	}

	counter := 0

	for l := range c.Numbers {
		var checked int

		for _, number := range c.Numbers[l] {
			if number.Checked {
				checked++
			}
		}

		if checked == 5 {
			counter++
		}
	}

	if !c.Completions.Horizontal.Check(counter) {
		return false
	}

	if !c.main.Completions.Horizontal.Add() {
		return false
	}

	return true
}

func (c *Card) IsVertical() bool {
	c.Completions.Vertical.Update(&c.main.Completions.Vertical)

	if c.main.Completions.Vertical.GetRemaining() == 0 {
		return false
	}

	if c.Completions.Vertical.GetRemaining() == 0 {
		return false
	}

	counter := 0

	for col := 0; col < 5; col++ {
		checked := 0

		for l := 0; l < 5; l++ {
			if c.Numbers[l][col].Checked {
				checked++
			}
		}

		if checked == 5 {
			counter++
		}
	}

	if !c.Completions.Vertical.Check(counter) {
		return false
	}

	if !c.main.Completions.Vertical.Add() {
		return false
	}

	return true
}

func (c *Card) IsChecked(drawedNumber int) bool {
	for _, line := range c.Numbers {
		for _, cardNumber := range line {
			if cardNumber.Number == drawedNumber && cardNumber.Checked {
				return true
			}
		}
	}

	return false
}

func (c *Card) SetConn(conn *websocket.Conn) bool {
	if conn == nil {
		return false
	}

	c.conn = conn

	return true
}

func (c *Card) SetNextRound(round int) bool {
	if c.NextRound == 0 && c.Round != round {
		c.NextRound = round
		c.UpdateCard()
		return true
	}

	return false
}

func (c *Card) ShortNumbers(numbers []Number) []Number {
	sort.Slice(numbers, func(i, j int) bool { return numbers[i].Number < numbers[j].Number })

	return numbers
}

func (c *Card) ToggleAutoplay() {
	if c.Autoplay {
		c.Autoplay = false
	} else {
		c.Autoplay = true
	}
}

func (c *Card) ToggleNumber(number int) bool {
	if c.CheckNumber(number) {
		return true
	}

	if c.UncheckNumber(number) {
		return true
	}

	return false
}

func (c *Card) UncheckNumber(number int) bool {
	for l, line := range c.Numbers {
		for col, column := range line {
			if column.Number == number && column.Checked {
				c.Numbers[l][col].Checked = false
				c.Checked--
				c.IsBingo()
				c.UpdateCard()
				return true
			}
		}
	}

	return false
}

func (c *Card) UpdateCard() error {
	if c.conn == nil {
		return nil
	}

	err := c.conn.WriteJSON(c)

	if err != nil {
		return err
	}

	return nil
}

// NewCard creates a new bingo card.
func NewCard(round *Round) Card {
	main, err := round.GetCard(0)

	card := Card{
		Autoplay: true,
		Card:     len(round.Cards) + 1,
		Round:    round.Round,
		Type:     round.Type,
	}

	if err == nil {
		card.main = main
		card.Completions = main.Completions
	} else {
		card.Completions = NewDefaultCompletions()
	}

	card.DrawCard()

	return card
}
