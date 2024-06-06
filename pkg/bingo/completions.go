package bingo

type Completions struct {
	Diagonal   Completion
	Full       Completion
	Horizontal Completion
	Total      Completion
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
		c.Full.Add()
	case "diagonal":
		c.Diagonal.Add()
	case "horizontal":
		c.Horizontal.Add()
	case "vertical":
		c.Vertical.Add()
	}
}

func NewDefaultCompletions() Completions {
	return Completions{
		Diagonal:   NewDefaultCompletion(1),
		Full:       NewDefaultCompletion(1),
		Horizontal: NewDefaultCompletion(1),
		Total:      NewDefaultCompletion(4),
		Vertical:   NewDefaultCompletion(1),
	}
}
