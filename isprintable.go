package piscine

func IsPrintable(str string) bool {
	runes := []rune(str)
	for i := range runes {
		if runes[i] < 32 {
			return false
		}
	}
	return true
}
