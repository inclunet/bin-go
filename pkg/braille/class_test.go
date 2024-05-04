package braille

import "testing"

func TestClass_Draw(t *testing.T) {
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
			if !HasOneOfThesWords(tt.c.Draw(), tt.want) {
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
