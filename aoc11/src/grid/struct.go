package grid

type Grid struct {
	Board               [10][10]int
	Flashes             int
	coordinateModifiers [8][2]int
}
