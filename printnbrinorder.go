package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var nb []int
	if n == 0 {
		nb = append(nb, 0)
	}
	for ; n > 0; n /= 10 {
		nb = append(nb, n%10)
	}
	SortIntegerTable(nb)
	for i := range nb {
		z01.PrintRune(rune('0' + nb[i]))
	}
}
