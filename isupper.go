package piscine

func IsUpper(str string) bool {
	runes := []rune(str)
	for i := range runes {
		if runes[i] < 'A' || runes[i] > 'Z' {
			return false
		}
	}
	return true
}
