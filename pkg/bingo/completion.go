package bingo

type Completion struct {
	Enabled  bool
	Max      int
	Quantity int
}

func (c *Completion) Check(i int) bool {
	if !c.Enabled {
		return false
	}

	if i > c.Max || i > 4 {
		return false
	}

	if i == c.Quantity {
		return false
	}

	if i > c.Quantity {
		c.Quantity = i
		return true
	}

	return false
}
