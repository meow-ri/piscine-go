package piscine

func AlphaCount(str string) int {
	runes := []rune(str)
	var n int = StrLen(str)
	var res int = 0
	for i := 0; i < n; i++ {
		if (rune[i] <= 'Z' && rune >= 'A') || (rune[i] <= 'z' && rune >= 'a') {
			res++
		}
	}
	return res
}
