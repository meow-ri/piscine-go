package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	runes := []rune(s)
	var res int = 0
	var neg bool = false
	if runes[0] == '-' {
		neg = true
		runes[0] = '0'
	} else if runes[0] == '+' {
		runes[0] = '0'
	}
	for _, r := range runes {
		res *= 10
		if r == '0' {
			res += 0
		} else if r == '1' {
			res += 1
		} else if r == '2' {
			res += 2
		} else if r == '3' {
			res += 3
		} else if r == '4' {
			res += 4
		} else if r == '5' {
			res += 5
		} else if r == '6' {
			res += 6
		} else if r == '7' {
			res += 7
		} else if r == '8' {
			res += 8
		} else if r == '9' {
			res += 9
		} else {
			return 0
		}
	}
	if neg {
		res *= -1
	}
	return res
}
