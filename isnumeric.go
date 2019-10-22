package piscine

func IsNumeric(str string) bool {
	runes := []rune(str)
	for i := range runes {
		if !IsRuneDigit(runes[i]) {
			return false
		}
	}
	return true
}
