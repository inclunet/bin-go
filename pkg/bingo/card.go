package bingo

import (
	"sort"

	"github.com/gorilla/websocket"
)

type BingoTypes struct {
	Diagonal   bool
	FullCard   bool
	Horizontal bool
	Vertical   bool
}

type Number struct {
	Checked bool
	Column  int
	Number  int
}

type Card struct {
	DisallowBingoTypes BingoTypes
	Autoplay           bool
	Bingo              bool
	Round              int
	Card               int
	conn               *websocket.Conn
	Checked            int
	LastNumber         int
	NextRound          int
	Type               int
	Numbers            [][5]Number
}

func (c *Card) CheckDrawedNumbers(mainCard Card) int {
	counter := 0

	if !c.Autoplay {
		return counter
	}

	if mainCard.Checked == 0 {
		return 0
	}

	for _, line := range mainCard.Numbers {
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
	if c.IsChecked(number) {
		return false
	}

	for l, line := range c.Numbers {
		for col, column := range line {
			if column.Number == number && !column.Checked {
				c.LastNumber = number
				c.Numbers[l][col].Checked = true
				c.Checked++
				c.Bingo = c.IsBingo()
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

	newNumber := GetRandomNumber(1, c.Type)

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
	drawNumber := Number{Checked: false, Column: column, Number: GetRandomNumber(columnStart, columnEnd)}

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

func (c *Card) IsBingo() bool {
	if c.Card == 1 {
		return false
	}

	if c.IsHorizontalCardBingo() {
		return true
	}

	if c.IsVerticalCardBingo() {
		return true
	}

	if c.IsDiagonalCardBingo() {
		return true
	}
	if c.IsFullCardBingo() {
		return true
	}

	return false
}

func (c *Card) IsDiagonalCardBingo() bool {
	if c.DisallowBingoTypes.Diagonal {
		return false
	}

	var checkedL int
	var checkedR int

	left := 0
	right := 4

	for l, line := range c.Numbers {
		if l <= 4 {
			if number := line[left]; number.Checked {
				checkedL++
			}

			if Number := line[right]; Number.Checked {
				checkedR++
			}

			left++
			right--
		}
	}

	if checkedL == 5 || checkedR == 5 {
		c.DisallowBingoTypes.Diagonal = true
		return true
	}

	return false
}

func (c *Card) IsFullCardBingo() bool {
	if !c.DisallowBingoTypes.FullCard {
		return c.Checked >= 24
	}

	return false
}

func (c *Card) IsHorizontalCardBingo() bool {
	if c.DisallowBingoTypes.Horizontal {
		return false
	}

	for _, line := range c.Numbers {
		var checked int

		for _, number := range line {
			if number.Checked {
				checked++
			}
		}

		if checked == 5 {
			c.DisallowBingoTypes.Horizontal = true
			return true
		}
	}

	return false
}

func (c *Card) IsVerticalCardBingo() bool {
	if c.DisallowBingoTypes.Vertical {
		return false
	}

	var checked [5]int

	for _, line := range c.Numbers {
		for col, number := range line {
			if number.Checked {
				checked[col]++
			}
		}
	}

	for _, number := range checked {
		if number == 5 {
			c.DisallowBingoTypes.Vertical = true
			return true
		}
	}

	return false
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

func (c *Card) ToggleNumber(mainCard Card, number int) bool {
	if c.IsChecked(number) && c.Card == 1 {
		return c.UncheckNumber(number)
	} else {
		if !mainCard.IsChecked(number) && c.Card > 1 {
			return false
		}

		return c.CheckNumber(number)
	}
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
func NewCard(round Round) (card Card) {
	card.Autoplay = true
	card.Card = len(round.Cards) + 1
	card.Round = round.Round
	card.Type = round.Type
	card.DrawCard()
	return card
}
