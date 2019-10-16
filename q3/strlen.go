package strlen

func StrLen(str string) int {
	runes := []rune(str)
	var count int = 0
	for value := range runes {
		count++
		value = value
	}
	return count
}
