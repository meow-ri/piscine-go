package strlen

func StrRev(str string) string {
	runes := []rune(str)
	var len int = 0
	var rRune rune
	for value := range runes {
		len++
		value = value
	}
	for i := 0; i < len/2; i++ {
		rRune = runes[i]
		runes[i] = runes[len-i-1]
		runes[len-i-1] = rRune
	}
	return string(runes)
}
