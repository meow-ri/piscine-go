package piscine

func AtoiBase(s string, base string) int {
	len1 := StrLen(s)
	len2 := StrLen(base)
	runes1 := []rune(s)
	runes2 := []rune(base)
	res := 0
	exists := false

	if len1 == 0 || len2 < 2 {
		return 0
	}
	for i := 0; i < len2; i++ {
		if runes2[i] == '-' || runes2[i] == '+' {
			return 0
		}
		for j := i + 1; j < len2; j++ {
			if runes2[i] == runes2[j] {
				return 0
			}
		}
	}
	if runes1[0] == '-' {
		return 0
	}
	pow := RecursivePower(len2, len1-1)
	for i := 0; i < len1; i++ {
		exists = false
		for j := 0; j < len2; j++ {
			if runes1[i] == runes2[j] {
				exists = true
				res = res + j*pow
				pow /= len2
				j = len2
			}
		}
		if !exists {
			return 0
		}
	}
	return res
}
