package piscine

func IsLower(str string) bool {
	runes := []rune(str)
	for i := range runes {
		if runes[i] < 'a' || runes[i] > 'z' {
			return false
		}
	}
	return true
}
