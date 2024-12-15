package bingo

import "fmt"

type Completions struct {
	Diagonal     Completion
	Full         Completion
	Horizontal   Completion
	Intermediary Completion
	Total        Completion
	Vertical     Completion
}

func (c *Completions) Update(completions *Completions) error {
	if err := c.Total.Update(&completions.Total); err != nil {
		return fmt.Errorf("error updating total completions %w", err)
	}

	if err := c.Intermediary.Update(&completions.Intermediary); err != nil {
		return fmt.Errorf("error updating intermediary completions %w", err)
	}

	if err := c.Full.Update(&completions.Full); err != nil {
		return fmt.Errorf("error updating full completions %w", err)
	}

	if err := c.Diagonal.Update(&completions.Diagonal); err != nil {
		return fmt.Errorf("error updating diagonal completions %w", err)
	}

	if err := c.Horizontal.Update(&completions.Horizontal); err != nil {
		return fmt.Errorf("error updating horizontal completions %w", err)
	}

	if err := c.Vertical.Update(&completions.Vertical); err != nil {
		return fmt.Errorf("error updating vertical completions %w", err)
	}

	return nil
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

func NewDefaultCompletions() *Completions {
	return &Completions{
		Diagonal:     NewDefaultCompletion(4),
		Full:         NewDefaultCompletion(1),
		Horizontal:   NewDefaultCompletion(4),
		Intermediary: NewDefaultCompletion(3),
		Total:        NewDefaultCompletion(4),
		Vertical:     NewDefaultCompletion(4),
	}
}
