package translator

import (
	"strconv"
)

type BrailleCode [8]bool
type BrailleCell [3][2]bool

var BrailleTable [256]string = [256]string{
	0:  " ",
	1:  "a",
	2:  ",",
	3:  "b",
	4:  ".",
	5:  "k",
	6:  ";",
	7:  "l",
	8:  "^",
	9:  "c",
	10: "i",
	11: "f",
	12: "í",
	13: "m",
	14: "s",
	15: "p",
	16: "~",
	17: "e",
	18: ":",
	19: "h",
	20: "*",
	21: "o",
	22: "!",
	23: "r",
	24: "´",
	25: "d",
	26: "j",
	27: "g",
	28: "ã",
	29: "n",
	30: "t",
	31: "q",
	32: "'",
	33: "â",
	34: "?",
	35: "ê",
	36: "-",
	37: "u",
	38: "\"",
	39: "v",
	40: " ",
	41: "ì",
	42: "õ",
	43: "à",
	44: "ó",
	45: "x",
	46: "è",
	47: "ç",
	48: "$",
	49: "@",
	50: "/",
	51: "ü",
	52: "º",
	53: "z",
	54: "=",
	55: "á",
	56: "|",
	57: "ô",
	58: "w",
	59: "ñ",
	60: "#",
	61: "y",
	62: "ú",
	63: "é",
}

func BinToCell(b byte) BrailleCell {
	return CodeToCell(BinToCode(b))
}

func BinToChar(b byte) string {
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

func BrailleToString(b []BrailleCell) (s string) {
	for _, c := range b {
		s += CellToChar(c)
	}

	return s
}

func CellToChar(b BrailleCell) string {
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

func CharToCell(c string) (b BrailleCell) {
	return CodeToCell(CharToCode(c))
}

func CharToCode(c string) BrailleCode {
	return IntToCode(CharToInt(c))
}

func CharToInt(c string) (i int) {
	for i, l := range BrailleTable {
		if l == c {
			return BinToInt(byte(i))
		}
	}

	return 0
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

func CodeToChar(b BrailleCode) string {
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

func IntToChar(i int) string {
	return BrailleTable[IntToBin(i)]
}

func IntToCode(i int) (b BrailleCode) {
	c := IntToBin(i)

	for d := 7; d >= 0; d-- {
		b[d] = c&(1<<uint(d)) > 0
	}

	return b
}

func StringToBraille(s string) (b []BrailleCell) {
	for _, c := range s {
		b = append(b, CharToCell(string(c)))
	}

	return b
}
