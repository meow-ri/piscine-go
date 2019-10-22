package piscine

func Compare(a, b string) int {
	l1 := StrLen(a)
	l2 := StrLen(b)
	n := l1
	runes1 := []rune(a)
	runes2 := []rune(b)
	if l1 == 0 && l2 == 0 {
		return 0
	} else if l1 == 0 {
		return -1
	} else if l2 == 0 {
		return 1
	} else {
		if l1 > l2 {
			n = l2
		}
		for i := 0; i < n; i++ {
			if runes1[i] > runes2[i] {
				return 1
			} else if runes1[i] < runes2[i] {
				return -1
			}
		}
		if l1 == l2 {
			return 0
		} else if l1 > l2 {
			return 1
		}
		return -1
	}
}
