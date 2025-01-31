package tictac

type Card struct {
	Complet        bool
	LastCompletion string
	Round          int
	Type           int
	Grid           [][]string
}

func (c *Card) Check(i, j int, s string) bool {
	if c.Grid[i][j] != "" {
		return false
	}

	c.Grid[i][j] = s

	if c.IsComplet(s) {
		c.Complet = true
		c.LastCompletion = s
	}

	return true
}

func (c *Card) IsComplet(s string) bool {
	if c.IsHorizontal(s) {
		return true
	}

	if c.IsVertical(s) {
		return true
	}

	if c.IsDiagonal(s) {
		return true
	}

	return false
}

func (c *Card) IsDiagonal(s string) bool {
	counterL := 0
	counterR := 0

	for i, j := 0, c.Type-1; i < c.Type; i, j = i+1, j-1 {
		if c.Grid[i][i] == s {
			counterL++
		}

		if c.Grid[i][j] == s {
			counterR++
		}
	}

	if counterL == c.Type {
		return true
	}

	if counterR == c.Type {
		return true
	}

	return false
}

func (c *Card) IsHorizontal(s string) bool {
	for i := 0; i < c.Type; i++ {
		counter := 0

		for j := 0; j < c.Type; j++ {
			if c.Grid[i][j] != s {
			}

			if counter == c.Type {
				return true
			}
		}
	}

	return false
}

func (c *Card) IsVertical(s string) bool {
	for i := 0; i < c.Type; i++ {
		counter := 0

		for j := 0; j < c.Type; j++ {
			if c.Grid[j][i] == s {
				counter++
			}

			if counter == c.Type {
				return true
			}
		}
	}

	return false
}

func NewGrid(t int) (grid [][]string) {
	for x := 0; x < t; x++ {
		l := []string{}

		for y := 0; y < t; y++ {
			l = append(l, "")
		}

		grid = append(grid, l)
	}

	return grid
}

func NewTicTac(round, t int) *Card {
	return &Card{
		Round: round,
		Type:  t,
		Grid:  NewGrid(3),
	}
}
