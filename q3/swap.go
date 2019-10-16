package swap

func Swap(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}
