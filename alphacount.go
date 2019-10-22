package piscine

func AlphaCount(str string) int {

	var n int = StrLen(str)
	var res int = 0
	for i := 0; i < n; i++ {
		if (str[i] <= 'Z' && str >= 'A') || (str[i] <= 'z' && str >= 'a') {
			res++
		}
	}
	return res
}
