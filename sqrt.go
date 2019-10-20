package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	} else {
		n := 0
		for i := 0; i*i <= nb; i++ {
			n = i
		}
		if nb == n*n {
			return n
		} else {
			return 0
		}
	}
}
