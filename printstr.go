package printstr

import "github.com/01-edu/z01"

func PrintStr(str string) {
	runes := []rune(str)
	for v := range runes {
		z01.PrintRune(runes[v])
	}
}
