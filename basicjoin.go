package piscine

func BasicJoin(strs []string) string {
	res := ""
	for i := range strs {
		res = Concat(res, strs[i])
	}
	return res
}
