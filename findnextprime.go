package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else if nb < 2 {
		return 2
	} else {
		if nb%2 == 0 {
			nb++
		}
		if IsPrime(nb) {
			return nb
		} else {
			return FindNextPrime(nb + 2)
		}
	}
}
