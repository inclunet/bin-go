package classes

import (
	"reflect"
	"testing"

	"github.com/inclunet/bin-go/pkg/translator"
)

func TestNewChallenge(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want Challenge
	}{
		{
			name: "should return a Challenge",
			args: args{
				word: "test",
			},
			want: Challenge{
				Braille: []translator.BrailleCell{
					{
						{false, true},
						{true, true},
						{true, false},
					},
					{
						{true, false},
						{false, true},
						{false, false},
					},
					{
						{false, true},
						{true, false},
						{true, false},
					},
					{
						{false, true},
						{true, true},
						{true, false},
					},
				},
				Word: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChallenge(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChallenge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChallenge_Check(t *testing.T) {
	type args struct {
		challenge Challenge
	}
	tests := []struct {
		name string
		c    *Challenge
		args args
		want bool
	}{
		{
			name: "should return true",
			c: &Challenge{
				Braille: []translator.BrailleCell{
					{
						{false, true},
						{true, true},
						{true, false},
					},
					{
						{true, false},
						{false, true},
						{false, false},
					},
					{
						{false, true},
						{true, false},
						{true, false},
					},
					{
						{false, true},
						{true, true},
						{true, false},
					},
				},
				Word: "test",
			},
			args: args{
				challenge: Challenge{
					Braille: []translator.BrailleCell{
						{
							{false, true},
							{true, true},
							{true, false},
						},
						{
							{true, false},
							{false, true},
							{false, false},
						},
						{
							{false, true},
							{true, false},
							{true, false},
						},
						{
							{false, true},
							{true, true},
							{true, false},
						},
					},
					Word: "test",
				},
			},
			want: true,
		},
		{
			name: "should return false",
			c: &Challenge{
				Braille: []translator.BrailleCell{
					{
						{false, true},
						{true, true},
						{true, false},
					},
					{
						{true, false},
						{false, true},
						{false, false},
					},
					{
						{false, true},
						{true, false},
						{true, false},
					},
					{
						{false, true},
						{true, true},
						{true, false},
					},
				},
				Word: "test",
			},
			args: args{
				challenge: Challenge{
					Braille: []translator.BrailleCell{
						{
							{false, true},
							{true, true},
							{true, false},
						},
						{
							{true, false},
							{false, true},
							{false, false},
						},
						{
							{false, true},
							{true, false},
							{true, false},
						},
						{
							{false, true},
							{true, true},
							{true, false},
						},
					},
					Word: "word",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Check(tt.args.challenge); got != tt.want {
				t.Errorf("Challenge.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
