package grid

import (
	"fmt"
)

type Grid struct {
	Board               [10][10]int
	Flashes             int
	coordinateModifiers [8][2]int
}

func (g *Grid) getMoves(current [2]int) [][2]int {
	moves := make([][2]int, 0)
	var newRow, newCol int

	for _, modifier := range g.coordinateModifiers {
		newRow = current[0] + modifier[0]
		newCol = current[1] + modifier[1]

		if g.validCoordinates(newRow, newCol) {
			moves = append(moves, [2]int{newRow, newCol})
		}
	}
	return moves
}

func (g *Grid) validCoordinates(row, col int) bool {
	return row < len(g.Board) && col < len(g.Board[0]) && row >= 0 && col >= 0
}

func (g *Grid) Display() {
	for _, r := range g.Board {
		for _, c := range r {
			fmt.Print(c)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println("Flashes:", g.Flashes)
}
