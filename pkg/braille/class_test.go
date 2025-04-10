package braille

import (
	"testing"
)

func TestClass_DrawWord(t *testing.T) {
	tests := []struct {
		name string
		c    *Class
		want []string
	}{
		{
			name: "should return a string",
			c: &Class{
				Description:         "This is a test class",
				RequiredPunctuation: 3,
				Rounds:              5,
				Words:               []string{"test", "word", "class"},
			},
			want: []string{"test", "word", "class"},
		},
		{
			name: "should return a string",
			c: &Class{
				Description:         "This is a test class",
				RequiredPunctuation: 3,
				Rounds:              5,
				Words:               []string{"test", "word", "class", "blue"},
			},
			want: []string{"test", "word", "class", "blue"},
		},
		{
			name: "should return \"\"",
			c: &Class{
				Words: []string{},
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !HasOneOfThesWords(tt.c.DrawWord(), tt.want) {
				t.Error("Class.Draw() = want one of given words")
			}
		})
	}
}

func HasOneOfThesWords(w string, words []string) bool {
	for _, word := range words {
		if w == word {
			return true
		}
	}
	return false
}

func TestClass_DrawAllowedChallenge(t *testing.T) {
	tests := []struct {
		name string
		c    *Class
		want []string
	}{
		{
			name: "should return a string",
			c: &Class{
				Description:         "This is a test class",
				RequiredPunctuation: 3,
				Rounds:              5,
				AllowedChallenges:   []string{"test", "word", "class"},
			},
			want: []string{"test", "word", "class"},
		},
		{
			name: "should return a string",
			c: &Class{
				Description:         "This is a test class",
				RequiredPunctuation: 3,
				Rounds:              5,
				AllowedChallenges:   []string{"test", "word", "class", "blue"},
			},
			want: []string{"test", "word", "class", "blue"},
		},
		{
			name: "should return \"\"",
			c: &Class{
				AllowedChallenges: []string{},
			},
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.DrawAllowedChallenge(); !HasOneOfThesWords(got, tt.want) {
				t.Errorf("Class.DrawAllowedChallenge() = %v, want %v", got, tt.want)
			}
		})
	}
}
