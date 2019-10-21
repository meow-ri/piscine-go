package piscine

import "github.com/01-edu/z01"

func QueensRec(nb int, pos [8]int) {
	if nb == 8 {
		for i := 0; i < 8; i++ {
			z01.PrintRune(rune('1' + pos[i]))
		}
		z01.PrintRune('\n')
	} else {
		for i := 0; i < 8; i++ {
			if IsSafe(nb, i, pos) {
				pos[nb] = i
				QueensRec(nb+1, pos)
			}
		}
	}
}
