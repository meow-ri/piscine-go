package piscine

func IsSafe(qn int, rp int, pos [8]int) bool {
	for i := 0; i < qn; i++ {
		t := pos[i]
		if t == rp || t == rp-(qn-i) || t == rp+(qn-i) {
			return false
		}
	}
	return true
}
