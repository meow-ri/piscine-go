package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var iRune rune = '0'
	var jRune rune = '0'
	var kRune rune = '0'
	var lRune rune = '0'
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				for l := 0; l <= 9; l++ {
					if i*10+j < k*10+l {
						z01.PrintRune(iRune)
						z01.PrintRune(jRune)
						z01.PrintRune(' ')
						z01.PrintRune(kRune)
						z01.PrintRune(lRune)
						if i != 9 || j != 8 || k != 9 || l != 9 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
					if lRune == '9' {
						lRune = '0'
					} else {
						lRune++
					}
				}
				if kRune == '9' {
					kRune = '0'
				} else {
					kRune++
				}
			}
			if jRune == '9' {
				jRune = '0'
			} else {
				jRune++
			}
		}
		if iRune == '9' {
			iRune = '0'
		} else {
			iRune++
		}
	}
	z01.PrintRune('\n')
}
