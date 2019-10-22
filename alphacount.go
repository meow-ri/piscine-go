package piscine

func AlphaCount(str string) int {
	runes := []rune(str)
	var n int = StrLen(str)
	var res int = 0
	for i := 0; i < n; i++ {
		if (runes[i] <= 'Z' && runes[i] >= 'A') || (runes[i] <= 'z' && runes[i] >= 'a') {
			res++
		}
	}
	return res
}
