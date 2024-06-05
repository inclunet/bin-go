package bingo

type Completions struct {
	Diagonal   Completion
	Full       Completion
	Horizontal Completion
	Vertical   Completion
}

func (c *Completions) Update(completions *Completions) {
	c.Full.Update(&completions.Full)
	c.Diagonal.Update(&completions.Diagonal)
	c.Horizontal.Update(&completions.Horizontal)
	c.Vertical.Update(&completions.Vertical)
}

func (c *Completions) Register(completion string) {
	switch completion {
	case "full":
		c.Full.AddCheck()
	case "diagonal":
		c.Diagonal.AddCheck()
	case "horizontal":
		c.Horizontal.AddCheck()
	case "vertical":
		c.Vertical.AddCheck()
	}
}

func NewDefaultCompletions() Completions {
	return Completions{
		Diagonal:   NewDefaultCompletion(1),
		Full:       NewDefaultCompletion(1),
		Horizontal: NewDefaultCompletion(1),
		Vertical:   NewDefaultCompletion(1),
	}
}
