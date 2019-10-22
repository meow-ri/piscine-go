package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	isFirstLetter := true
	for i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' && isFirstLetter {
			runes[i] -= 32
			isFirstLetter = false
		} else if (IsRuneDigit(runes[i]) || (runes[i] >= 'A' && runes[i] <= 'Z')) && isFirstLetter {
			isFirstLetter = false
		} else if !isFirstLetter && runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		} else if !isFirstLetter && !IsRuneDigit(runes[i]) && !IsRuneAlpha(runes[i]) {
			isFirstLetter = true
		}
	}
	return string(runes)
}
