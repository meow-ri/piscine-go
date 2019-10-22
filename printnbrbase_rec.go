package piscine

import "github.com/01-edu/z01"

func PrintNbrBaseRec(n int, runes []rune, len int) {
	if n/len != 0 {
		PrintNbrBaseRec(n/len, runes, len)
	}
	mod := n % len
	if mod < 0 {
		mod = -mod
	}
	z01.PrintRune(runes[mod])
}
