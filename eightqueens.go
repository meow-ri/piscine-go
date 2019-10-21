package piscine

func EightQueens() {
	pos := [8]int{0, 0, 0, 0, 0, 0, 0}
	QueensRec(0, pos)
}
