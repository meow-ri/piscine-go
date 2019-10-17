package piscine

func SortIntegerTable(table []int) {
	var len int = 0
	for _, v := range table {
		v = v
		len++
	}
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if table[i] > table[j] {
				Swap(&table[i], &table[j])
			}
		}
	}
}
