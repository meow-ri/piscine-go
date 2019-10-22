package piscine

func NRune(s string, n int) rune {
	len := StrLen(s)
	if len >= n && len != 0 && n > 0 {
		runes := []rune(s)
		return runes[n-1]
	} else {
		return 0
	}
}
