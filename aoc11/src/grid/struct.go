package grid

type Grid struct {
	board [10][10]int
}

func NewGrid(board [10][10]int) Grid {
	grid := Grid{board: board}
	return grid
}
