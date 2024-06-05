package bingo

type Completion struct {
	Enabled  bool
	Max      int
	Quantity int
}

func (c *Completion) AddCheck() {
	if c.Check(c.Quantity + 1) {
		if c.GetRemaining() == 0 {
			c.Enabled = false
		}
	}
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

func (c *Completion) GetRemaining() int {
	if i := c.Max - c.Quantity; i >= 0 {
		return i
	}

	return 0
}

func (c *Completion) Update(completion *Completion) {
	c.Enabled = completion.Enabled

	if c.Quantity <= completion.GetRemaining() {
		c.Max = completion.GetRemaining()
	}
}

func NewDefaultCompletion(max int) Completion {
	return Completion{
		Enabled:  max > 0,
		Max:      max,
		Quantity: 0,
	}
}
