package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i := range runes {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		}
	}
	return string(runes)
}
