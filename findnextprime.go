package piscine

func FindNextPrime(nb int) int {

	if IsPrime(nb) {
		return nb
	} else {
		for i := 1; i <= nb; i++ {
			nb += i
			if IsPrime(nb) {
				return nb
			}
		}
	}
}
