package piscine

func IsAlpha(str string) bool {
	runes := []rune(str)
	for i := range runes {
		if !IsRuneAlpha(runes[i]) && !IsRuneDigit(runes[i]) {
			return false
		}
	}
	return true
}
