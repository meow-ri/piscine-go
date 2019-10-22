package piscine

func Join(strs []string, sep string) string {
	res := ""
	for i := range strs {
		if i != 0 {
			res = Concat(res, sep)
		}
		res = Concat(res, strs[i])
	}
	return res
}
