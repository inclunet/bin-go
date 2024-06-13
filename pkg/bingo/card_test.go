package bingo

import (
	"testing"
)

func GenerateFakeDiagonalCardBingo(i int) *Card {
	card, err := GenerateFakeRound().AddCard()

	if err != nil {
		return nil
	}

	for l, r := 0, 4; l < 5; l, r = l+1, r-1 {
		if i == 0 {
			card.Numbers[l][l].Checked = true
		} else {
			card.Numbers[l][r].Checked = true
		}
	}

	return card
}

func GenerateFakeFullCardBingo() *Card {
	card, err := GenerateFakeRound().AddCard()

	if err != nil {
		return nil
	}

	for l := 0; l < 5; l++ {
		for c := 0; c < 5; c++ {
			card.Numbers[l][c].Checked = true
		}
	}

	return card
}

func GenerateFakeHorizontalCardBingo(max int, i ...int) *Card {
	card, err := GenerateFakeRound().AddCard()

	if err != nil {
		return nil
	}

	card.main.Completions.Horizontal = NewDefaultCompletion(max)
	card.Completions.Horizontal = NewDefaultCompletion(max)

	for _, l := range i {
		for c := 0; c < 5; c++ {
			card.Numbers[l][c].Checked = true
		}
	}

	return card
}

func GenerateFakeVerticalCardBingo(max int, i ...int) *Card {
	card, err := GenerateFakeRound().AddCard()

	if err != nil {
		return nil
	}

	card.main.Completions.Vertical = NewDefaultCompletion(max)
	card.Completions.Vertical = NewDefaultCompletion(max)

	for _, c := range i {
		for l := 0; l < 5; l++ {
			card.Numbers[l][c].Checked = true
		}
	}

	return card
}

func GenerateFakeNewCard() *Card {
	card, err := GenerateFakeRound().AddCard()

	if err != nil {
		return nil
	}

	return card
}

func GenerateFakeRound() *Round {
	round := NewRound(&Bingo{}, 75)

	return &round
}

func TestCard_IsDiagonal(t *testing.T) {
	tests := []struct {
		name string
		c    *Card
		want bool
	}{

		{
			name: "Left Diagonal completed",
			c:    GenerateFakeDiagonalCardBingo(0),
			want: true,
		},
		{
			name: "Right Diagonal completed",
			c:    GenerateFakeDiagonalCardBingo(1),
			want: true,
		},
		{
			name: "New Card",
			c:    GenerateFakeNewCard(),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsDiagonal(); got != tt.want {
				t.Errorf("Card.IsDiagonal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_IsHorizontal(t *testing.T) {
	tests := []struct {
		name string
		c    *Card
		want bool
	}{

		{
			name: "Horizontal 0 completed with 0 remaining",
			c:    GenerateFakeHorizontalCardBingo(0, 0),
			want: false,
		},
		{
			name: "Horizontal 0 completed",
			c:    GenerateFakeHorizontalCardBingo(1, 0),
			want: true,
		},
		{
			name: "Horizontal 1 completed",
			c:    GenerateFakeHorizontalCardBingo(1, 1),
			want: true,
		},
		{
			name: "Horizontal 2 completed",
			c:    GenerateFakeHorizontalCardBingo(1, 2),
			want: true,
		},
		{
			name: "Horizontal 3 completed",
			c:    GenerateFakeHorizontalCardBingo(1, 3),
			want: true,
		},
		{
			name: "Horizontal 4 completed",
			c:    GenerateFakeHorizontalCardBingo(1, 4),
			want: true,
		},
		{
			name: "Horizontal 0 and 1 completed with max 1 allowed",
			c:    GenerateFakeHorizontalCardBingo(1, 0, 1),
			want: false,
		},
		{
			name: "Horizontal 0 completed with max 2 allowed",
			c:    GenerateFakeHorizontalCardBingo(2, 0),
			want: true,
		},
		{
			name: "Horizontal 0 and 1 completed with max 2 allowed",
			c:    GenerateFakeHorizontalCardBingo(2, 0, 1),
			want: true,
		},
		{
			name: "Horizontal 0, 1 and  3 completed with max 2 allowed",
			c:    GenerateFakeHorizontalCardBingo(2, 0, 1, 3),
			want: false,
		},
		{
			name: "New Card",
			c:    GenerateFakeNewCard(),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsHorizontal(); got != tt.want {
				t.Errorf("Card.IsHorizontal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_IsVertical(t *testing.T) {
	tests := []struct {
		name string
		c    *Card
		want bool
	}{
		{
			name: "Vertical 0 completed with 0 remaining",
			c:    GenerateFakeVerticalCardBingo(0, 0),
			want: false,
		},
		{
			name: "Vertical 0 completed",
			c:    GenerateFakeVerticalCardBingo(1, 0),
			want: true,
		},
		{
			name: "Vertical 1 completed",
			c:    GenerateFakeVerticalCardBingo(1, 1),
			want: true,
		},
		{
			name: "Vertical 2 completed",
			c:    GenerateFakeVerticalCardBingo(1, 2),
			want: true,
		},
		{
			name: "Vertical 3 completed",
			c:    GenerateFakeVerticalCardBingo(1, 3),
			want: true,
		},
		{
			name: "Vertical 4 completed",
			c:    GenerateFakeVerticalCardBingo(1, 4),
			want: true,
		},
		{
			name: "Vertical 0 and 1 completed with max 1 allowed",
			c:    GenerateFakeVerticalCardBingo(1, 0, 1),
			want: false,
		},
		{
			name: "Vertical 0 completed with max 2 allowed",
			c:    GenerateFakeVerticalCardBingo(2, 0),
			want: true,
		},
		{
			name: "Vertical 0 and 1 completed with max 2 allowed",
			c:    GenerateFakeVerticalCardBingo(2, 0, 1),
			want: true,
		},
		{
			name: "Vertical 0, 1 and  3 completed with max 2 allowed",
			c:    GenerateFakeVerticalCardBingo(2, 0, 1, 3),
			want: false,
		},
		{
			name: "New Card",
			c:    GenerateFakeNewCard(),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsVertical(); got != tt.want {
				t.Errorf("Card.IsVertical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_IsBingo(t *testing.T) {
	tests := []struct {
		name string
		c    *Card
		want bool
	}{
		{
			name: "Full Card Completed shows Bingo",
			c:    GenerateFakeFullCardBingo(),
			want: true,
		},
		{
			name: "New Card",
			c:    GenerateFakeNewCard(),
			want: false,
		},
		{
			name: "Horizontal 0 completed",
			c:    GenerateFakeHorizontalCardBingo(1, 0),
			want: true,
		},
		{
			name: "Horizontal 0 and 1 completed with max 1 allowed",
			c:    GenerateFakeHorizontalCardBingo(1, 0, 1),
			want: false,
		},
		{
			name: "Vertical 0 completed",
			c:    GenerateFakeVerticalCardBingo(1, 0),
			want: true,
		},
		{
			name: "Vertical 0 and 1 completed with max 1 allowed",
			c:    GenerateFakeVerticalCardBingo(1, 0, 1),
			want: false,
		},
		{
			name: "Left Diagonal completed",
			c:    GenerateFakeDiagonalCardBingo(0),
			want: true,
		},
		{
			name: "Right Diagonal completed",
			c:    GenerateFakeDiagonalCardBingo(1),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsBingo(); got != tt.want {
				t.Errorf("Card.IsBingo() = %v, want %v", got, tt.want)
			}
		})
	}
}
