package translator

import (
	"strconv"
)

type BrailleCode [8]bool
type BrailleCell [3][2]bool

var BrailleTable [256]rune = [256]rune{
	0:  32,
	1:  97,
	2:  44,
	3:  98,
	4:  46,
	5:  107,
	6:  59,
	7:  108,
	8:  94,
	9:  99,
	10: 105,
	11: 102,
	12: 237,
	13: 109,
	14: 115,
	15: 112,
	16: 126,
	17: 101,
	18: 58,
	19: 104,
	20: 42,
	21: 111,
	22: 33,
	23: 114,
	24: 180,
	25: 100,
	26: 106,
	27: 103,
	28: 227,
	29: 110,
	30: 116,
	31: 113,
	32: 39,
	33: 225,
	34: 63,
	35: 234,
	36: 45,
	37: 117,
	38: 34,
	39: 118,
	40: 40,
	41: 236,
	42: 245,
	43: 224,
	44: 243,
	45: 120,
	46: 232,
	47: 231,
	48: 36,
	49: 64,
	50: 47,
	51: 252,
	52: 186,
	53: 122,
	54: 61,
	55: 225,
	56: 124,
	57: 244,
	58: 119,
	59: 241,
	60: 35,
	61: 121,
	62: 250,
	63: 233,
}

func BinToCell(b byte) BrailleCell {
	return CodeToCell(BinToCode(b))
}

func BinToChar(b byte) rune {
	return IntToChar(BinToInt(b))
}

func BinToCode(b byte) (c BrailleCode) {
	return IntToCode(BinToInt(b))
}

func BinToInt(b byte) (i int) {
	for d := 0; d < 8; d++ {
		if b&(1<<uint(d)) > 0 {
			i = i*10 + d + 1
		}
	}

	return i
}

func BrailleToString(b []BrailleCell) string {
	var r []rune

	for _, c := range b {
		r = append(r, CellToChar(c))
	}

	return string(r)
}

func CellToChar(b BrailleCell) rune {
	return IntToChar(CellToInt(b))
}

func CellToCode(b BrailleCell) (c BrailleCode) {
	return IntToCode(CellToInt(b))
}

func CellToInt(b BrailleCell) (i int) {
	code := ""

	if b[0][0] {
		code += "1"
	}

	if b[1][0] {
		code += "2"
	}

	if b[2][0] {
		code += "3"
	}

	if b[0][1] {
		code += "4"
	}

	if b[1][1] {
		code += "5"
	}

	if b[2][1] {
		code += "6"
	}

	i, _ = strconv.Atoi(code)

	return i
}

func CharToCell(c rune) (b BrailleCell) {
	return CodeToCell(CharToCode(c))
}

func CharToCode(c rune) BrailleCode {
	return IntToCode(GetRuneIndex(c))
}

func CodeToCell(b BrailleCode) (c BrailleCell) {
	c[0][0] = b[0]
	c[0][1] = b[3]
	c[1][0] = b[1]
	c[1][1] = b[4]
	c[2][0] = b[2]
	c[2][1] = b[5]

	return c
}

func CodeToChar(b BrailleCode) rune {
	return IntToChar(CodeToInt(b))
}

func CodeToInt(b BrailleCode) (i int) {
	return CellToInt(CodeToCell(b))
}

func IntToBin(i int) (b byte) {
	for h := 1; i > 0; h++ {
		d := i % 10
		p := byte(1 << uint(d-1))
		b |= p
		i = i / 10
	}

	return b
}

func IntToCell(i int) BrailleCell {
	return CodeToCell(IntToCode(i))
}

func IntToChar(i int) rune {
	return BrailleTable[IntToBin(i)]
}

func IntToCode(i int) (b BrailleCode) {
	c := IntToBin(i)

	for d := 7; d >= 0; d-- {
		b[d] = c&(1<<uint(d)) > 0
	}

	return b
}

// GetRuneIndex returns the index of a rune in the BrailleTable
// or the index of the rune in the Unicode Braille Patterns
// range if the rune is not in the BrailleTable.
func GetRuneIndex(r rune) int {
	if r >= 10240 && r <= 10495 {
		return int(r - 10240)
	}

	for i := range BrailleTable {
		if BrailleTable[i] == r {
			return i
		}
	}

	return 0
}

// ToggleRune returns the Braille representation of a rune
// if the rune is in the BrailleTable, or the Unicode Braille
// Patterns representation of the rune if the rune is not in
// the BrailleTable.
func ToggleRune(r rune) rune {
	i := GetRuneIndex(r)

	if r >= 10240 && r <= 10495 {
		if i >= 0 && i < len(BrailleTable) {
			return BrailleTable[i]
		}
	}

	return rune(i + 10240)
}

func StringToCells(s string) (b []BrailleCell) {
	for _, r := range s {
		b = append(b, CharToCell(r))
	}

	return b
}

func ToggleString(s string) string {
	var r []rune

	for _, c := range s {
		r = append(r, ToggleRune(c))
	}

	return string(r)
}
