package bingo

type Number struct {
	Checked bool
	Column  int
	Number  int
}

func IsDuplicatedNumber(numbers []Number, number Number) bool {
	for _, storedNumber := range numbers {
		if number.Number == storedNumber.Number {
			return true
		}
	}

	return false
}
