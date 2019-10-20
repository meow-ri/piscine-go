package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		var n int = 1
		for i := 1; i <= power; i++ {
			n *= nb
		}
		return n
	}
}
