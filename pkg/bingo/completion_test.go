package bingo

import "testing"

func TestCompletion_Check(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		c    *Completion
		args args
		want bool
	}{

		{
			name: "Max is 4, Quantity is 4, i is 4",
			c: &Completion{
				Max:      4,
				Quantity: 4,
			},
			args: args{
				i: 4,
			},
			want: false,
		},
		{
			name: "Max is 0, Quantity is 0, i is 0",
			c: &Completion{
				Max:      0,
				Quantity: 0,
			},
			args: args{
				i: 0,
			},
			want: false,
		},
		{
			name: "Max is 0, Quantity is 0, i is 1",
			c: &Completion{
				Max:      0,
				Quantity: 0,
			},
			args: args{
				i: 1,
			},
			want: false,
		},
		{
			name: "Max is 1, Quantity is 0, i is 0",
			c: &Completion{
				Max:      1,
				Quantity: 0,
			},
			args: args{
				i: 0,
			},
			want: false,
		},
		{
			name: "Max is 1, Quantity is 0, i is 1",
			c: &Completion{
				Max:      1,
				Quantity: 0,
			},
			args: args{
				i: 1,
			},
			want: true,
		},
		{
			name: "Max is 1, Quantity is 0, i is 2",
			c: &Completion{
				Max:      1,
				Quantity: 0,
			},
			args: args{
				i: 2,
			},
			want: false,
		},
		{
			name: "Max is 1, Quantity is 1, i is 0",
			c: &Completion{
				Max:      1,
				Quantity: 1,
			},
			args: args{
				i: 0,
			},
			want: false,
		},
		{
			name: "Max is 1, Quantity is 1, i is 1",
			c: &Completion{
				Max:      1,
				Quantity: 1,
			},
			args: args{
				i: 1,
			},
			want: false,
		},
		{
			name: "Max is 1, Quantity is 1, i is 2",
			c: &Completion{
				Max:      1,
				Quantity: 1,
			},
			args: args{
				i: 2,
			},
			want: false,
		},
		{
			name: "Max is 2, Quantity is 0, i is 0",
			c: &Completion{
				Max:      2,
				Quantity: 0,
			},
			args: args{
				i: 0,
			},
			want: false,
		},
		{
			name: "Max is 2, Quantity is 0, i is 1",
			c: &Completion{
				Max:      2,
				Quantity: 0,
			},
			args: args{
				i: 1,
			},
			want: true,
		},
		{
			name: "Max is 2, Quantity is 0, i is 2",
			c: &Completion{
				Max:      2,
				Quantity: 0,
			},
			args: args{
				i: 2,
			},
			want: true,
		},
		{
			name: "Max is 2, Quantity is 0, i is 3",
			c: &Completion{
				Max:      2,
				Quantity: 0,
			},
			args: args{
				i: 3,
			},
			want: false,
		},
		{
			name: "Max is 2, Quantity is 1, i is 0",
			c: &Completion{
				Max:      2,
				Quantity: 1,
			},
			args: args{
				i: 0,
			},
			want: false,
		},
		{
			name: "Max is 2, Quantity is 1, i is 1",
			c: &Completion{
				Max:      2,
				Quantity: 1,
			},
			args: args{
				i: 1,
			},
			want: false,
		},
		{
			name: "Max is 2, Quantity is 1, i is 2",
			c: &Completion{
				Max:      2,
				Quantity: 1,
			},
			args: args{
				i: 2,
			},
			want: true,
		},
		{
			name: "Max is 2, Quantity is 1, i is 3",
			c: &Completion{
				Max:      2,
				Quantity: 1,
			},
			args: args{
				i: 3,
			},
			want: false,
		},
		{
			name: "Max is 2, Quantity is 2, i is 0",
			c: &Completion{
				Max:      2,
				Quantity: 2,
			},
			args: args{
				i: 0,
			},
			want: false,
		},
		{
			name: "Max is 2, Quantity is 2, i is 1",
			c: &Completion{
				Max:      2,
				Quantity: 2,
			},
			args: args{
				i: 1,
			},
			want: false,
		},
		{
			name: "Max is 2, Quantity is 2, i is 2",
			c: &Completion{
				Max:      2,
				Quantity: 2,
			},
			args: args{
				i: 2,
			},
			want: false,
		},
		{
			name: "Max is 2, Quantity is 2, i is 3",
			c: &Completion{
				Max:      2,
				Quantity: 2,
			},
			args: args{
				i: 3,
			},
			want: false,
		},
		{
			name: "Max is 3, Quantity is 0, i is 0",
			c: &Completion{
				Max:      3,
				Quantity: 0,
			},
			args: args{
				i: 0,
			},
			want: false,
		},
		{
			name: "Max is 3, Quantity is 0, i is 1",
			c: &Completion{
				Max:      3,
				Quantity: 0,
			},
			args: args{
				i: 1,
			},
			want: true,
		},
		{
			name: "Max is 3, Quantity is 0, i is 2",
			c: &Completion{
				Max:      2,
				Quantity: 0,
			},
			args: args{
				i: 2,
			},
			want: true,
		},
		{
			name: "Max is 3, Quantity is 0, i is 3",
			c: &Completion{
				Max:      3,
				Quantity: 0,
			},
			args: args{
				i: 3,
			},
			want: true,
		},
		{
			name: "Max is 3, Quantity is 0, i is 4",
			c: &Completion{
				Max:      3,
				Quantity: 0,
			},
			args: args{
				i: 4,
			},
			want: false,
		},
		{
			name: "Max is 3, Quantity is 1, i is 0",
			c: &Completion{
				Max:      3,
				Quantity: 1,
			},
			args: args{
				i: 0,
			},
			want: false,
		},
		{
			name: "Max is 3, Quantity is 1, i is 1",
			c: &Completion{
				Max:      3,
				Quantity: 1,
			},
			args: args{
				i: 1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Check(tt.args.i); got != tt.want {
				t.Errorf("Completion.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
