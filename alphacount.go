package piscine

func AlphaCount(str string) int {
	var res int = 0
	for _, i := range str {
		if (i <= 'Z' && i >= 'A') || (i <= 'z' && i >= 'a') {
			res++
		}
	}
	return res
}
