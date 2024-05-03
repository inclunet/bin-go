package translator

import (
	"reflect"
	"testing"
)

func TestCellToCode(t *testing.T) {
	type args struct {
		b BrailleCell
	}
	tests := []struct {
		name  string
		args  args
		wantC BrailleCode
	}{
		{
			name:  "get space",
			args:  args{b: BrailleCell{{false, false}, {false, false}, {false, false}}},
			wantC: BrailleCode{false, false, false, false, false, false},
		},
		{
			name:  "get a",
			args:  args{b: BrailleCell{{true, false}, {false, false}, {false, false}}},
			wantC: BrailleCode{true, false, false, false, false, false},
		},
		{
			name:  "get b",
			args:  args{b: BrailleCell{{true, false}, {true, false}, {false, false}}},
			wantC: BrailleCode{true, true, false, false, false, false},
		},
		{
			name:  "get c",
			args:  args{b: BrailleCell{{true, true}, {false, false}, {false, false}}},
			wantC: BrailleCode{true, false, false, true, false, false},
		},
		{
			name:  "get d",
			args:  args{b: BrailleCell{{true, true}, {false, true}, {false, false}}},
			wantC: BrailleCode{true, false, false, true, true, false},
		},
		{
			name:  "get e",
			args:  args{b: BrailleCell{{true, false}, {false, true}, {false, false}}},
			wantC: BrailleCode{true, false, false, false, true, false},
		},
		{
			name:  "get f",
			args:  args{b: BrailleCell{{true, true}, {true, false}, {false, false}}},
			wantC: BrailleCode{true, true, false, true, false, false},
		},
		{
			name:  "get g",
			args:  args{b: BrailleCell{{true, true}, {true, true}, {false, false}}},
			wantC: BrailleCode{true, true, false, true, true, false},
		},
		{
			name:  "get h",
			args:  args{b: BrailleCell{{true, false}, {true, true}, {false, false}}},
			wantC: BrailleCode{true, true, false, false, true, false},
		},
		{
			name:  "get i",
			args:  args{b: BrailleCell{{false, true}, {true, false}, {false, false}}},
			wantC: BrailleCode{false, true, false, true, false, false},
		},
		{
			name:  "get j",
			args:  args{b: BrailleCell{{false, true}, {true, true}, {false, false}}},
			wantC: BrailleCode{false, true, false, true, true, false},
		},
		{
			name:  "get k",
			args:  args{b: BrailleCell{{true, false}, {false, false}, {true, false}}},
			wantC: BrailleCode{true, false, true, false, false, false},
		},
		{
			name:  "get l",
			args:  args{b: BrailleCell{{true, false}, {true, false}, {true, false}}},
			wantC: BrailleCode{true, true, true, false, false, false},
		},
		{
			name:  "get m",
			args:  args{b: BrailleCell{{true, true}, {false, false}, {true, false}}},
			wantC: BrailleCode{true, false, true, true, false, false},
		},
		{
			name:  "get n",
			args:  args{b: BrailleCell{{true, true}, {false, true}, {true, false}}},
			wantC: BrailleCode{true, false, true, true, true, false},
		},
		{
			name:  "get o",
			args:  args{b: BrailleCell{{true, false}, {false, true}, {true, false}}},
			wantC: BrailleCode{true, false, true, false, true, false},
		},
		{
			name:  "get p",
			args:  args{b: BrailleCell{{true, true}, {true, false}, {true, false}}},
			wantC: BrailleCode{true, true, true, true, false, false},
		},
		{
			name:  "get q",
			args:  args{b: BrailleCell{{true, true}, {true, true}, {true, false}}},
			wantC: BrailleCode{true, true, true, true, true, false},
		},
		{
			name:  "get r",
			args:  args{b: BrailleCell{{true, false}, {true, true}, {true, false}}},
			wantC: BrailleCode{true, true, true, false, true, false},
		},
		{
			name:  "get s",
			args:  args{b: BrailleCell{{false, true}, {true, false}, {true, false}}},
			wantC: BrailleCode{false, true, true, true, false, false},
		},
		{
			name:  "get t",
			args:  args{b: BrailleCell{{false, true}, {true, true}, {true, false}}},
			wantC: BrailleCode{false, true, true, true, true, false},
		},
		{
			name:  "get u",
			args:  args{b: BrailleCell{{true, false}, {false, false}, {true, true}}},
			wantC: BrailleCode{true, false, true, false, false, true},
		},
		{
			name:  "get v",
			args:  args{b: BrailleCell{{true, false}, {true, false}, {true, true}}},
			wantC: BrailleCode{true, true, true, false, false, true},
		},
		{
			name:  "get x",
			args:  args{b: BrailleCell{{true, true}, {false, false}, {true, true}}},
			wantC: BrailleCode{true, false, true, true, false, true},
		},
		{
			name:  "get y",
			args:  args{b: BrailleCell{{true, true}, {false, true}, {true, true}}},
			wantC: BrailleCode{true, false, true, true, true, true},
		},
		{
			name:  "get z",
			args:  args{b: BrailleCell{{true, false}, {false, true}, {true, true}}},
			wantC: BrailleCode{true, false, true, false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := CellToCode(tt.args.b); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("CellToCode() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestCellToInt(t *testing.T) {
	type args struct {
		b BrailleCell
	}
	tests := []struct {
		name  string
		args  args
		wantI int
	}{
		{
			name:  "get space",
			args:  args{b: BrailleCell{{false, false}, {false, false}, {false, false}}},
			wantI: 0,
		},
		{
			name:  "get a",
			args:  args{b: BrailleCell{{true, false}, {false, false}, {false, false}}},
			wantI: 1,
		},
		{
			name:  "get b",
			args:  args{b: BrailleCell{{true, false}, {true, false}, {false, false}}},
			wantI: 12,
		},
		{
			name:  "get c",
			args:  args{b: BrailleCell{{true, true}, {false, false}, {false, false}}},
			wantI: 14,
		},
		{
			name:  "get d",
			args:  args{b: BrailleCell{{true, true}, {false, true}, {false, false}}},
			wantI: 145,
		},
		{
			name:  "get e",
			args:  args{b: BrailleCell{{true, false}, {false, true}, {false, false}}},
			wantI: 15,
		},
		{
			name:  "get f",
			args:  args{b: BrailleCell{{true, true}, {true, false}, {false, false}}},
			wantI: 124,
		},
		{
			name:  "get g",
			args:  args{b: BrailleCell{{true, true}, {true, true}, {false, false}}},
			wantI: 1245,
		},
		{
			name:  "get h",
			args:  args{b: BrailleCell{{true, false}, {true, true}, {false, false}}},
			wantI: 125,
		},
		{
			name:  "get i",
			args:  args{b: BrailleCell{{false, true}, {true, false}, {false, false}}},
			wantI: 24,
		},
		{
			name:  "get j",
			args:  args{b: BrailleCell{{false, true}, {true, true}, {false, false}}},
			wantI: 245,
		},
		{
			name:  "get k",
			args:  args{b: BrailleCell{{true, false}, {false, false}, {true, false}}},
			wantI: 13,
		},
		{
			name:  "get l",
			args:  args{b: BrailleCell{{true, false}, {true, false}, {true, false}}},
			wantI: 123,
		},
		{
			name:  "get m",
			args:  args{b: BrailleCell{{true, true}, {false, false}, {true, false}}},
			wantI: 134,
		},
		{
			name:  "get n",
			args:  args{b: BrailleCell{{true, false}, {false, true}, {true, false}}},
			wantI: 135,
		},
		{
			name:  "get o",
			args:  args{b: BrailleCell{{true, false}, {false, true}, {true, false}}},
			wantI: 135,
		},
		{
			name:  "get p",
			args:  args{b: BrailleCell{{true, true}, {true, false}, {true, false}}},
			wantI: 1234,
		},
		{
			name:  "get q",
			args:  args{b: BrailleCell{{true, true}, {true, true}, {true, false}}},
			wantI: 12345,
		},
		{
			name:  "get r",
			args:  args{b: BrailleCell{{true, false}, {true, true}, {true, false}}},
			wantI: 1235,
		},
		{
			name:  "get s",
			args:  args{b: BrailleCell{{false, true}, {true, false}, {true, false}}},
			wantI: 234,
		},
		{
			name:  "get t",
			args:  args{b: BrailleCell{{false, true}, {true, true}, {true, false}}},
			wantI: 2345,
		},
		{
			name:  "get u",
			args:  args{b: BrailleCell{{true, false}, {false, false}, {true, true}}},
			wantI: 136,
		},
		{
			name:  "get v",
			args:  args{b: BrailleCell{{true, false}, {true, false}, {true, true}}},
			wantI: 1236,
		},
		{
			name:  "get x",
			args:  args{b: BrailleCell{{true, true}, {false, false}, {true, true}}},
			wantI: 1346,
		},
		{
			name:  "get y",
			args:  args{b: BrailleCell{{true, false}, {false, true}, {true, true}}},
			wantI: 1356,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := CellToInt(tt.args.b); gotI != tt.wantI {
				t.Errorf("CellToInt() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestCharToCell(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name  string
		args  args
		wantB BrailleCell
	}{
		{
			name:  "get space",
			args:  args{c: " "},
			wantB: BrailleCell{{false, false}, {false, false}, {false, false}},
		},
		{
			name:  "get a",
			args:  args{c: "a"},
			wantB: BrailleCell{{true, false}, {false, false}, {false, false}},
		},
		{
			name:  "get b",
			args:  args{c: "b"},
			wantB: BrailleCell{{true, false}, {true, false}, {false, false}},
		},
		{
			name:  "get c",
			args:  args{c: "c"},
			wantB: BrailleCell{{true, true}, {false, false}, {false, false}},
		},
		{
			name:  "get d",
			args:  args{c: "d"},
			wantB: BrailleCell{{true, true}, {false, true}, {false, false}},
		},
		{
			name:  "get e",
			args:  args{c: "e"},
			wantB: BrailleCell{{true, false}, {false, true}, {false, false}},
		},
		{
			name:  "get f",
			args:  args{c: "f"},
			wantB: BrailleCell{{true, true}, {true, false}, {false, false}},
		},
		{
			name:  "get g",
			args:  args{c: "g"},
			wantB: BrailleCell{{true, true}, {true, true}, {false, false}},
		},
		{
			name:  "get h",
			args:  args{c: "h"},
			wantB: BrailleCell{{true, false}, {true, true}, {false, false}},
		},
		{
			name:  "get i",
			args:  args{c: "i"},
			wantB: BrailleCell{{false, true}, {true, false}, {false, false}},
		},
		{
			name:  "get j",
			args:  args{c: "j"},
			wantB: BrailleCell{{false, true}, {true, true}, {false, false}},
		},
		{
			name:  "get k",
			args:  args{c: "k"},
			wantB: BrailleCell{{true, false}, {false, false}, {true, false}},
		},
		{
			name:  "get l",
			args:  args{c: "l"},
			wantB: BrailleCell{{true, false}, {true, false}, {true, false}},
		},
		{
			name:  "get m",
			args:  args{c: "m"},
			wantB: BrailleCell{{true, true}, {false, false}, {true, false}},
		},
		{
			name:  "get n",
			args:  args{c: "n"},
			wantB: BrailleCell{{true, true}, {false, true}, {true, false}},
		},
		{
			name:  "get o",
			args:  args{c: "o"},
			wantB: BrailleCell{{true, false}, {false, true}, {true, false}},
		},
		{
			name:  "get p",
			args:  args{c: "p"},
			wantB: BrailleCell{{true, true}, {true, false}, {true, false}},
		},
		{
			name:  "get q",
			args:  args{c: "q"},
			wantB: BrailleCell{{true, true}, {true, true}, {true, false}},
		},
		{
			name:  "get r",
			args:  args{c: "r"},
			wantB: BrailleCell{{true, false}, {true, true}, {true, false}},
		},
		{
			name:  "get s",
			args:  args{c: "s"},
			wantB: BrailleCell{{false, true}, {true, false}, {true, false}},
		},
		{
			name:  "get t",
			args:  args{c: "t"},
			wantB: BrailleCell{{false, true}, {true, true}, {true, false}},
		},
		{
			name:  "get u",
			args:  args{c: "u"},
			wantB: BrailleCell{{true, false}, {false, false}, {true, true}},
		},
		{
			name:  "get v",
			args:  args{c: "v"},
			wantB: BrailleCell{{true, false}, {true, false}, {true, true}},
		},
		{
			name:  "get x",
			args:  args{c: "x"},
			wantB: BrailleCell{{true, true}, {false, false}, {true, true}},
		},
		{
			name:  "get y",
			args:  args{c: "y"},
			wantB: BrailleCell{{true, true}, {false, true}, {true, true}},
		},
		{
			name:  "get z",
			args:  args{c: "z"},
			wantB: BrailleCell{{true, false}, {false, true}, {true, true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := CharToCell(tt.args.c); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("CharToCell() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestCharToCode(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name string
		args args
		want BrailleCode
	}{
		{
			name: "get space",
			args: args{c: " "},
			want: BrailleCode{false, false, false, false, false, false},
		},
		{
			name: "get a",
			args: args{c: "a"},
			want: BrailleCode{true, false, false, false, false, false},
		},
		{
			name: "get b",
			args: args{c: "b"},
			want: BrailleCode{true, true, false, false, false, false},
		},
		{
			name: "get c",
			args: args{c: "c"},
			want: BrailleCode{true, false, false, true, false, false},
		},
		{
			name: "get d",
			args: args{c: "d"},
			want: BrailleCode{true, false, false, true, true, false},
		},
		{
			name: "get e",
			args: args{c: "e"},
			want: BrailleCode{true, false, false, false, true, false},
		},
		{
			name: "get f",
			args: args{c: "f"},
			want: BrailleCode{true, true, false, true, false, false},
		},
		{
			name: "get g",
			args: args{c: "g"},
			want: BrailleCode{true, true, false, true, true, false},
		},
		{
			name: "get h",
			args: args{c: "h"},
			want: BrailleCode{true, true, false, false, true, false},
		},
		{
			name: "get i",
			args: args{c: "i"},
			want: BrailleCode{false, true, false, true, false, false},
		},
		{
			name: "get j",
			args: args{c: "j"},
			want: BrailleCode{false, true, false, true, true, false},
		},
		{
			name: "get k",
			args: args{c: "k"},
			want: BrailleCode{true, false, true, false, false, false},
		},
		{
			name: "get l",
			args: args{c: "l"},
			want: BrailleCode{true, true, true, false, false, false},
		},
		{
			name: "get m",
			args: args{c: "m"},
			want: BrailleCode{true, false, true, true, false, false},
		},
		{
			name: "get n",
			args: args{c: "n"},
			want: BrailleCode{true, false, true, true, true, false},
		},
		{
			name: "get o",
			args: args{c: "o"},
			want: BrailleCode{true, false, true, false, true, false},
		},
		{
			name: "get p",
			args: args{c: "p"},
			want: BrailleCode{true, true, true, true, false, false},
		},
		{
			name: "get q",
			args: args{c: "q"},
			want: BrailleCode{true, true, true, true, true, false},
		},
		{
			name: "get r",
			args: args{c: "r"},
			want: BrailleCode{true, true, true, false, true, false},
		},
		{
			name: "get s",
			args: args{c: "s"},
			want: BrailleCode{false, true, true, true, false, false},
		},
		{
			name: "get t",
			args: args{c: "t"},
			want: BrailleCode{false, true, true, true, true, false},
		},
		{
			name: "get u",
			args: args{c: "u"},
			want: BrailleCode{true, false, true, false, false, true},
		},
		{
			name: "get v",
			args: args{c: "v"},
			want: BrailleCode{true, true, true, false, false, true},
		},
		{
			name: "get x",
			args: args{c: "x"},
			want: BrailleCode{true, false, true, true, false, true},
		},
		{
			name: "get y",
			args: args{c: "y"},
			want: BrailleCode{true, false, true, true, true, true},
		},
		{
			name: "get z",
			args: args{c: "z"},
			want: BrailleCode{true, false, true, false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CharToCode(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CharToCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharToInt(t *testing.T) {
	type args struct {
		c string
	}
	tests := []struct {
		name  string
		args  args
		wantI int
	}{
		{
			name:  "get space",
			args:  args{c: " "},
			wantI: 0,
		},
		{
			name:  "get a",
			args:  args{c: "a"},
			wantI: 1,
		},
		{
			name:  "get b",
			args:  args{c: "b"},
			wantI: 12,
		},
		{
			name:  "get c",
			args:  args{c: "c"},
			wantI: 14,
		},
		{
			name:  "get d",
			args:  args{c: "d"},
			wantI: 145,
		},
		{
			name:  "get e",
			args:  args{c: "e"},
			wantI: 15,
		},
		{
			name:  "get f",
			args:  args{c: "f"},
			wantI: 124,
		},
		{
			name:  "get g",
			args:  args{c: "g"},
			wantI: 1245,
		},
		{
			name:  "get h",
			args:  args{c: "h"},
			wantI: 125,
		},
		{
			name:  "get i",
			args:  args{c: "i"},
			wantI: 24,
		},
		{
			name:  "get j",
			args:  args{c: "j"},
			wantI: 245,
		},
		{
			name:  "get k",
			args:  args{c: "k"},
			wantI: 13,
		},
		{
			name:  "get l",
			args:  args{c: "l"},
			wantI: 123,
		},
		{
			name:  "get m",
			args:  args{c: "m"},
			wantI: 134,
		},
		{
			name:  "get n",
			args:  args{c: "n"},
			wantI: 1345,
		},
		{
			name:  "get o",
			args:  args{c: "o"},
			wantI: 135,
		},
		{
			name:  "get p",
			args:  args{c: "p"},
			wantI: 1234,
		},
		{
			name:  "get q",
			args:  args{c: "q"},
			wantI: 12345,
		},
		{
			name:  "get r",
			args:  args{c: "r"},
			wantI: 1235,
		},
		{
			name:  "get s",
			args:  args{c: "s"},
			wantI: 234,
		},
		{
			name:  "get t",
			args:  args{c: "t"},
			wantI: 2345,
		},
		{
			name:  "get u",
			args:  args{c: "u"},
			wantI: 136,
		},
		{
			name:  "get v",
			args:  args{c: "v"},
			wantI: 1236,
		},
		{
			name:  "get x",
			args:  args{c: "x"},
			wantI: 1346,
		},
		{
			name:  "get y",
			args:  args{c: "y"},
			wantI: 13456,
		},
		{
			name:  "get z",
			args:  args{c: "z"},
			wantI: 1356,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := CharToInt(tt.args.c); gotI != tt.wantI {
				t.Errorf("CharToInt() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestCodeToCell(t *testing.T) {
	type args struct {
		b BrailleCode
	}
	tests := []struct {
		name  string
		args  args
		wantC BrailleCell
	}{
		{
			name:  "get space",
			args:  args{b: BrailleCode{false, false, false, false, false, false}},
			wantC: BrailleCell{{false, false}, {false, false}, {false, false}},
		},
		{
			name:  "get a",
			args:  args{b: BrailleCode{true, false, false, false, false, false}},
			wantC: BrailleCell{{true, false}, {false, false}, {false, false}},
		},
		{
			name:  "get b",
			args:  args{b: BrailleCode{true, true, false, false, false, false}},
			wantC: BrailleCell{{true, false}, {true, false}, {false, false}},
		},
		{
			name:  "get c",
			args:  args{b: BrailleCode{true, false, false, true, false, false}},
			wantC: BrailleCell{{true, true}, {false, false}, {false, false}},
		},
		{
			name:  "get d",
			args:  args{b: BrailleCode{true, false, false, true, true, false}},
			wantC: BrailleCell{{true, true}, {false, true}, {false, false}},
		},
		{
			name:  "get e",
			args:  args{b: BrailleCode{true, false, false, false, true, false}},
			wantC: BrailleCell{{true, false}, {false, true}, {false, false}},
		},
		{
			name:  "get f",
			args:  args{b: BrailleCode{true, true, false, true, false, false}},
			wantC: BrailleCell{{true, true}, {true, false}, {false, false}},
		},
		{
			name:  "get g",
			args:  args{b: BrailleCode{true, true, false, true, true, false}},
			wantC: BrailleCell{{true, true}, {true, true}, {false, false}},
		},
		{
			name:  "get h",
			args:  args{b: BrailleCode{true, true, false, false, true, false}},
			wantC: BrailleCell{{true, false}, {true, true}, {false, false}},
		},
		{
			name:  "get i",
			args:  args{b: BrailleCode{false, true, false, true, false, false}},
			wantC: BrailleCell{{false, true}, {true, false}, {false, false}},
		},
		{
			name:  "get j",
			args:  args{b: BrailleCode{false, true, false, true, true, false}},
			wantC: BrailleCell{{false, true}, {true, true}, {false, false}},
		},
		{
			name:  "get k",
			args:  args{b: BrailleCode{true, false, true, false, false, false}},
			wantC: BrailleCell{{true, false}, {false, false}, {true, false}},
		},
		{
			name:  "get l",
			args:  args{b: BrailleCode{true, true, true, false, false, false}},
			wantC: BrailleCell{{true, false}, {true, false}, {true, false}},
		},
		{
			name:  "get m",
			args:  args{b: BrailleCode{true, false, true, true, false, false}},
			wantC: BrailleCell{{true, true}, {false, false}, {true, false}},
		},
		{
			name:  "get n",
			args:  args{b: BrailleCode{true, false, true, true, true, false}},
			wantC: BrailleCell{{true, true}, {false, true}, {true, false}},
		},
		{
			name:  "get o",
			args:  args{b: BrailleCode{true, false, true, false, true, false}},
			wantC: BrailleCell{{true, false}, {false, true}, {true, false}},
		},
		{
			name:  "get p",
			args:  args{b: BrailleCode{true, true, true, true, false, false}},
			wantC: BrailleCell{{true, true}, {true, false}, {true, false}},
		},
		{
			name:  "get q",
			args:  args{b: BrailleCode{true, true, true, true, true, false}},
			wantC: BrailleCell{{true, true}, {true, true}, {true, false}},
		},
		{
			name:  "get r",
			args:  args{b: BrailleCode{true, true, true, false, true, false}},
			wantC: BrailleCell{{true, false}, {true, true}, {true, false}},
		},
		{
			name:  "get s",
			args:  args{b: BrailleCode{false, true, true, true, false, false}},
			wantC: BrailleCell{{false, true}, {true, false}, {true, false}},
		},
		{
			name:  "get t",
			args:  args{b: BrailleCode{false, true, true, true, true, false}},
			wantC: BrailleCell{{false, true}, {true, true}, {true, false}},
		},
		{
			name:  "get u",
			args:  args{b: BrailleCode{true, false, true, false, false, true}},
			wantC: BrailleCell{{true, false}, {false, false}, {true, true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := CodeToCell(tt.args.b); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("BrailleToCell() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestCodeToChar(t *testing.T) {
	type args struct {
		b BrailleCode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "get space",
			args: args{b: BrailleCode{false, false, false, false, false, false}},
			want: " ",
		},
		{
			name: "get a",
			args: args{b: BrailleCode{true, false, false, false, false, false}},
			want: "a",
		},
		{
			name: "get b",
			args: args{b: BrailleCode{true, true, false, false, false, false}},
			want: "b",
		},
		{
			name: "get c",
			args: args{b: BrailleCode{true, false, false, true, false, false}},
			want: "c",
		},
		{
			name: "get d",
			args: args{b: BrailleCode{true, false, false, true, true, false}},
			want: "d",
		},
		{
			name: "get e",
			args: args{b: BrailleCode{true, false, false, false, true, false}},
			want: "e",
		},
		{
			name: "get f",
			args: args{b: BrailleCode{true, true, false, true, false, false}},
			want: "f",
		},
		{
			name: "get g",
			args: args{b: BrailleCode{true, true, false, true, true, false}},
			want: "g",
		},
		{
			name: "get h",
			args: args{b: BrailleCode{true, true, false, false, true, false}},
			want: "h",
		},
		{
			name: "get i",
			args: args{b: BrailleCode{false, true, false, true, false, false}},
			want: "i",
		},
		{
			name: "get j",
			args: args{b: BrailleCode{false, true, false, true, true, false}},
			want: "j",
		},
		{
			name: "get k",
			args: args{b: BrailleCode{true, false, true, false, false, false}},
			want: "k",
		},
		{
			name: "get l",
			args: args{b: BrailleCode{true, true, true, false, false, false}},
			want: "l",
		},
		{
			name: "get m",
			args: args{b: BrailleCode{true, false, true, true, false, false}},
			want: "m",
		},
		{
			name: "get n",
			args: args{b: BrailleCode{true, false, true, true, true, false}},
			want: "n",
		},
		{
			name: "get o",
			args: args{b: BrailleCode{true, false, true, false, true, false}},
			want: "o",
		},
		{
			name: "get p",
			args: args{b: BrailleCode{true, true, true, true, false, false}},
			want: "p",
		},
		{
			name: "get q",
			args: args{b: BrailleCode{true, true, true, true, true, false}},
			want: "q",
		},
		{
			name: "get r",
			args: args{b: BrailleCode{true, true, true, false, true, false}},
			want: "r",
		},
		{
			name: "get s",
			args: args{b: BrailleCode{false, true, true, true, false, false}},
			want: "s",
		},
		{
			name: "get t",
			args: args{b: BrailleCode{false, true, true, true, true, false}},
			want: "t",
		},
		{
			name: "get u",
			args: args{b: BrailleCode{true, false, true, false, false, true}},
			want: "u",
		},
		{
			name: "get v",
			args: args{b: BrailleCode{true, true, true, false, false, true}},
			want: "v",
		},
		{
			name: "get x",
			args: args{b: BrailleCode{true, false, true, true, false, true}},
			want: "x",
		},
		{
			name: "get y",
			args: args{b: BrailleCode{true, false, true, true, true, true}},
			want: "y",
		},
		{
			name: "get z",
			args: args{b: BrailleCode{true, false, true, false, true, true}},
			want: "z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CodeToChar(tt.args.b); got != tt.want {
				t.Errorf("CodeToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodeToInt(t *testing.T) {
	type args struct {
		b BrailleCode
	}
	tests := []struct {
		name  string
		args  args
		wantI int
	}{
		{
			name:  "get space",
			args:  args{b: BrailleCode{false, false, false, false, false, false}},
			wantI: 0,
		},
		{
			name:  "get a",
			args:  args{b: BrailleCode{true, false, false, false, false, false}},
			wantI: 1,
		},
		{
			name:  "get b",
			args:  args{b: BrailleCode{true, true, false, false, false, false}},
			wantI: 12,
		},
		{
			name:  "get c",
			args:  args{b: BrailleCode{true, false, false, true, false, false}},
			wantI: 14,
		},
		{
			name:  "get d",
			args:  args{b: BrailleCode{true, false, false, true, true, false}},
			wantI: 145,
		},
		{
			name:  "get e",
			args:  args{b: BrailleCode{true, false, false, false, true, false}},
			wantI: 15,
		},
		{
			name:  "get f",
			args:  args{b: BrailleCode{true, true, false, true, false, false}},
			wantI: 124,
		},
		{
			name:  "get g",
			args:  args{b: BrailleCode{true, true, false, true, true, false}},
			wantI: 1245,
		},
		{
			name:  "get h",
			args:  args{b: BrailleCode{true, true, false, false, true, false}},
			wantI: 125,
		},
		{
			name:  "get i",
			args:  args{b: BrailleCode{false, true, false, true, false, false}},
			wantI: 24,
		},
		{
			name:  "get j",
			args:  args{b: BrailleCode{false, true, false, true, true, false}},
			wantI: 245,
		},
		{
			name:  "get k",
			args:  args{b: BrailleCode{true, false, true, false, false, false}},
			wantI: 13,
		},
		{
			name:  "get l",
			args:  args{b: BrailleCode{true, true, true, false, false, false}},
			wantI: 123,
		},
		{
			name:  "get m",
			args:  args{b: BrailleCode{true, false, true, true, false, false}},
			wantI: 134,
		},
		{
			name:  "get n",
			args:  args{b: BrailleCode{true, false, true, true, true, false}},
			wantI: 1345,
		},
		{
			name:  "get o",
			args:  args{b: BrailleCode{true, false, true, false, true, false}},
			wantI: 135,
		},
		{
			name:  "get p",
			args:  args{b: BrailleCode{true, true, true, true, false, false}},
			wantI: 1234,
		},
		{
			name:  "get q",
			args:  args{b: BrailleCode{true, true, true, true, true, false}},
			wantI: 12345,
		},
		{
			name:  "get r",
			args:  args{b: BrailleCode{true, true, true, false, true, false}},
			wantI: 1235,
		},
		{
			name:  "get s",
			args:  args{b: BrailleCode{false, true, true, true, false, false}},
			wantI: 234,
		},
		{
			name:  "get t",
			args:  args{b: BrailleCode{false, true, true, true, true, false}},
			wantI: 2345,
		},
		{
			name:  "get u",
			args:  args{b: BrailleCode{true, false, true, false, false, true}},
			wantI: 136,
		},
		{
			name:  "get v",
			args:  args{b: BrailleCode{true, true, true, false, false, true}},
			wantI: 1236,
		},
		{
			name:  "get x",
			args:  args{b: BrailleCode{true, false, true, true, false, true}},
			wantI: 1346,
		},
		{
			name:  "get y",
			args:  args{b: BrailleCode{true, false, true, true, true, true}},
			wantI: 13456,
		},
		{
			name:  "get z",
			args:  args{b: BrailleCode{true, false, true, false, true, true}},
			wantI: 1356,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := CodeToInt(tt.args.b); gotI != tt.wantI {
				t.Errorf("CodeToInt() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestIntToCell(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want BrailleCell
	}{

		{
			name: "get space",
			args: args{i: 0},
			want: BrailleCell{{false, false}, {false, false}, {false, false}},
		},
		{
			name: "get a",
			args: args{i: 1},
			want: BrailleCell{{true, false}, {false, false}, {false, false}},
		},
		{
			name: "get b",
			args: args{i: 12},
			want: BrailleCell{{true, false}, {true, false}, {false, false}},
		},
		{
			name: "get c",
			args: args{i: 14},
			want: BrailleCell{{true, true}, {false, false}, {false, false}},
		},
		{
			name: "get d",
			args: args{i: 145},
			want: BrailleCell{{true, true}, {false, true}, {false, false}},
		},
		{
			name: "get e",
			args: args{i: 15},
			want: BrailleCell{{true, false}, {false, true}, {false, false}},
		},
		{
			name: "get f",
			args: args{i: 124},
			want: BrailleCell{{true, true}, {true, false}, {false, false}},
		},
		{
			name: "get g",
			args: args{i: 1245},
			want: BrailleCell{{true, true}, {true, true}, {false, false}},
		},
		{
			name: "get h",
			args: args{i: 125},
			want: BrailleCell{{true, false}, {true, true}, {false, false}},
		},
		{
			name: "get i",
			args: args{i: 24},
			want: BrailleCell{{false, true}, {true, false}, {false, false}},
		},
		{
			name: "get j",
			args: args{i: 245},
			want: BrailleCell{{false, true}, {true, true}, {false, false}},
		},
		{
			name: "get k",
			args: args{i: 13},
			want: BrailleCell{{true, false}, {false, false}, {true, false}},
		},
		{
			name: "get l",
			args: args{i: 123},
			want: BrailleCell{{true, false}, {true, false}, {true, false}},
		},
		{
			name: "get m",
			args: args{i: 134},
			want: BrailleCell{{true, true}, {false, false}, {true, false}},
		},
		{
			name: "get n",
			args: args{i: 1345},
			want: BrailleCell{{true, true}, {false, true}, {true, false}},
		},
		{
			name: "get o",
			args: args{i: 135},
			want: BrailleCell{{true, false}, {false, true}, {true, false}},
		},
		{
			name: "get p",
			args: args{i: 1234},
			want: BrailleCell{{true, true}, {true, false}, {true, false}},
		},
		{
			name: "get q",
			args: args{i: 12345},
			want: BrailleCell{{true, true}, {true, true}, {true, false}},
		},
		{
			name: "get r",
			args: args{i: 1235},
			want: BrailleCell{{true, false}, {true, true}, {true, false}},
		},
		{
			name: "get s",
			args: args{i: 234},
			want: BrailleCell{{false, true}, {true, false}, {true, false}},
		},
		{
			name: "get t",
			args: args{i: 2345},
			want: BrailleCell{{false, true}, {true, true}, {true, false}},
		},
		{
			name: "get u",
			args: args{i: 136},
			want: BrailleCell{{true, false}, {false, false}, {true, true}},
		},
		{
			name: "get v",
			args: args{i: 1236},
			want: BrailleCell{{true, false}, {true, false}, {true, true}},
		},
		{
			name: "get x",
			args: args{i: 1346},
			want: BrailleCell{{true, true}, {false, false}, {true, true}},
		},
		{
			name: "get y",
			args: args{i: 13456},
			want: BrailleCell{{true, true}, {false, true}, {true, true}},
		},
		{
			name: "get z",
			args: args{i: 1356},
			want: BrailleCell{{true, false}, {false, true}, {true, true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToCell(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntToCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToChar(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{
			name: "get space",
			args: args{i: 0},
			want: " ",
		},
		{
			name: "get a",
			args: args{i: 1},
			want: "a",
		},
		{
			name: "get b",
			args: args{i: 12},
			want: "b",
		},
		{
			name: "get c",
			args: args{i: 14},
			want: "c",
		},
		{
			name: "get d",
			args: args{i: 145},
			want: "d",
		},
		{
			name: "get e",
			args: args{i: 15},
			want: "e",
		},
		{
			name: "get f",
			args: args{i: 124},
			want: "f",
		},
		{
			name: "get g",
			args: args{i: 1245},
			want: "g",
		},
		{
			name: "get h",
			args: args{i: 125},
			want: "h",
		},
		{
			name: "get i",
			args: args{i: 24},
			want: "i",
		},
		{
			name: "get j",
			args: args{i: 245},
			want: "j",
		},
		{
			name: "get k",
			args: args{i: 13},
			want: "k",
		},
		{
			name: "get l",
			args: args{i: 123},
			want: "l",
		},
		{
			name: "get m",
			args: args{i: 134},
			want: "m",
		},
		{
			name: "get n",
			args: args{i: 1345},
			want: "n",
		},
		{
			name: "get o",
			args: args{i: 135},
			want: "o",
		},
		{
			name: "get p",
			args: args{i: 1234},
			want: "p",
		},
		{
			name: "get q",
			args: args{i: 12345},
			want: "q",
		},
		{
			name: "get r",
			args: args{i: 1235},
			want: "r",
		},
		{
			name: "get s",
			args: args{i: 234},
			want: "s",
		},
		{
			name: "get t",
			args: args{i: 2345},
			want: "t",
		},
		{
			name: "get u",
			args: args{i: 136},
			want: "u",
		},
		{
			name: "get v",
			args: args{i: 1236},
			want: "v",
		},
		{
			name: "get x",
			args: args{i: 1346},
			want: "x",
		},
		{
			name: "get y",
			args: args{i: 13456},
			want: "y",
		},
		{
			name: "get z",
			args: args{i: 1356},
			want: "z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntToChar(tt.args.i); got != tt.want {
				t.Errorf("IntToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntToCode(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name  string
		args  args
		wantB BrailleCode
	}{

		{
			name:  "get space",
			args:  args{i: 0},
			wantB: BrailleCode{false, false, false, false, false, false, false, false},
		},
		{
			name:  "get a",
			args:  args{i: 1},
			wantB: BrailleCode{true, false, false, false, false, false, false, false},
		},
		{
			name:  "get b",
			args:  args{i: 12},
			wantB: BrailleCode{true, true, false, false, false, false, false, false},
		},
		{
			name:  "get c",
			args:  args{i: 14},
			wantB: BrailleCode{true, false, false, true, false, false, false, false},
		},
		{
			name:  "get d",
			args:  args{i: 145},
			wantB: BrailleCode{true, false, false, true, true, false, false, false},
		},
		{
			name:  "get e",
			args:  args{i: 15},
			wantB: BrailleCode{true, false, false, false, true, false, false, false},
		},
		{
			name:  "get f",
			args:  args{i: 124},
			wantB: BrailleCode{true, true, false, true, false, false, false, false},
		},
		{
			name:  "get g",
			args:  args{i: 1245},
			wantB: BrailleCode{true, true, false, true, true, false, false, false},
		},
		{
			name:  "get h",
			args:  args{i: 125},
			wantB: BrailleCode{true, true, false, false, true, false, false, false},
		},
		{
			name:  "get i",
			args:  args{i: 24},
			wantB: BrailleCode{false, true, false, true, false, false, false, false},
		},
		{
			name:  "get j",
			args:  args{i: 245},
			wantB: BrailleCode{false, true, false, true, true, false, false, false},
		},
		{
			name:  "get k",
			args:  args{i: 13},
			wantB: BrailleCode{true, false, true, false, false, false, false, false},
		},
		{
			name:  "get l",
			args:  args{i: 123},
			wantB: BrailleCode{true, true, true, false, false, false, false, false},
		},
		{
			name:  "get m",
			args:  args{i: 134},
			wantB: BrailleCode{true, false, true, true, false, false, false, false},
		},
		{
			name:  "get n",
			args:  args{i: 1345},
			wantB: BrailleCode{true, false, true, true, true, false, false, false},
		},
		{
			name:  "get o",
			args:  args{i: 135},
			wantB: BrailleCode{true, false, true, false, true, false, false, false},
		},
		{
			name:  "get p",
			args:  args{i: 1234},
			wantB: BrailleCode{true, true, true, true, false, false, false, false},
		},
		{
			name:  "get q",
			args:  args{i: 12345},
			wantB: BrailleCode{true, true, true, true, true, false, false, false},
		},
		{
			name:  "get r",
			args:  args{i: 1235},
			wantB: BrailleCode{true, true, true, false, true, false, false, false},
		},
		{
			name:  "get s",
			args:  args{i: 234},
			wantB: BrailleCode{false, true, true, true, false, false, false, false},
		},
		{
			name:  "get t",
			args:  args{i: 2345},
			wantB: BrailleCode{false, true, true, true, true, false, false, false},
		},
		{
			name:  "get u",
			args:  args{i: 136},
			wantB: BrailleCode{true, false, true, false, false, true, false, false},
		},
		{
			name:  "get v",
			args:  args{i: 1236},
			wantB: BrailleCode{true, true, true, false, false, true, false, false},
		},
		{
			name:  "get x",
			args:  args{i: 1346},
			wantB: BrailleCode{true, false, true, true, false, true, false, false},
		},
		{
			name:  "get y",
			args:  args{i: 13456},
			wantB: BrailleCode{true, false, true, true, true, true, false, false},
		},
		{
			name:  "get z",
			args:  args{i: 1356},
			wantB: BrailleCode{true, false, true, false, true, true, false, false},
		},
		{
			name:  "get ç",
			args:  args{i: 12346},
			wantB: BrailleCode{true, true, true, true, false, true, false, false},
		},
		{
			name:  "get é",
			args:  args{i: 123456},
			wantB: BrailleCode{true, true, true, true, true, true, false, false},
		},
		{
			name:  "get braille á",
			args:  args{i: 12356},
			wantB: BrailleCode{true, true, true, false, true, true, false, false},
		},
		{
			name:  "get braille <?>",
			args:  args{i: 2346},
			wantB: BrailleCode{false, true, true, true, false, true, false, false},
		},
		{
			name:  "get braille ú",
			args:  args{i: 23456},
			wantB: BrailleCode{false, true, true, true, true, true, false, false},
		},
		{
			name:  "get braille â",
			args:  args{i: 16},
			wantB: BrailleCode{true, false, false, false, false, true, false, false},
		},
		{
			name:  "get braille ê",
			args:  args{i: 126},
			wantB: BrailleCode{true, true, false, false, false, true, false, false},
		},
		{
			name:  "get braille <?>",
			args:  args{i: 146},
			wantB: BrailleCode{true, false, false, true, false, true, false, false},
		},
		{
			name:  "get braille ô",
			args:  args{i: 1456},
			wantB: BrailleCode{true, false, false, true, true, true, false, false},
		},
		{
			name:  "get braille @",
			args:  args{i: 156},
			wantB: BrailleCode{true, false, false, false, true, true, false, false},
		},
		{
			name:  "get braille à",
			args:  args{i: 1246},
			wantB: BrailleCode{true, true, false, true, false, true, false, false},
		},
		{
			name:  "get braille <?>",
			args:  args{i: 12456},
			wantB: BrailleCode{true, true, false, true, true, true, false, false},
		},
		{
			name:  "get braille <?>",
			args:  args{i: 1256},
			wantB: BrailleCode{true, true, false, false, true, true, false, false},
		},
		{
			name:  "get braille õ",
			args:  args{i: 246},
			wantB: BrailleCode{false, true, false, true, false, true, false, false},
		},
		{
			name:  "get braille w",
			args:  args{i: 2456},
			wantB: BrailleCode{false, true, false, true, true, true, false, false},
		},
		{
			name:  "get braille 2",
			args:  args{i: 2},
			wantB: BrailleCode{false, true, false, false, false, false, false, false},
		},
		{
			name:  "get braille 3",
			args:  args{i: 3},
			wantB: BrailleCode{false, false, true, false, false, false, false, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := IntToCode(tt.args.i); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("IntToBraille() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestCellToChar(t *testing.T) {
	type args struct {
		b BrailleCell
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CellToChar(tt.args.b); got != tt.want {
				t.Errorf("CellToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
