package piscine

func Index(s string, toFind string) int {
	l1 := StrLen(s)
	l2 := StrLen(toFind)
	found := true
	if l2 < 1 {
		return 0
	} else if l1 < 1 {
		return -1
	} else {
		runes1 := []rune(s)
		runes2 := []rune(toFind)
		for i := 0; i < l1; i++ {
			found = true
			for j := 0; j < l2; j++ {
				if runes1[i+j] != runes2[j] {
					found = false
					j = l2
				}
			}
			if found {
				return i
			}
		}
		return -1
	}
}
