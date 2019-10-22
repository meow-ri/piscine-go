package piscine

func LastRune(s string) rune {
	len := StrLen(s)
	return NRune(s, len)
}
