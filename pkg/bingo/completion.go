package bingo

type Completion struct {
	Enabled  bool
	Max      int
	Quantity int
}

func (c *Completion) Add() bool {
	return c.Check(c.Quantity + 1)
}

func (c *Completion) Check(i int) bool {
	if c.GetRemaining() == 0 {
		return false
	}

	if i == c.Quantity {
		return false
	}

	if i > c.Max {
		return false
	}

	if i <= c.Quantity {
		return false
	}

	c.Quantity = i

	return true
}

func (c *Completion) GetRemaining() int {
	if i := c.Max - c.Quantity; i >= 0 {
		return i
	}

	return 0
}

func (c *Completion) Update(completion *Completion) {
	if c.Quantity <= completion.GetRemaining() {
		c.Max = completion.GetRemaining()
	}
}

func NewDefaultCompletion(max int) Completion {
	return Completion{
		Max:      max,
		Quantity: 0,
	}
}
