package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a / *b
	d := *a % *b
	*a = c
	*b = d
}
