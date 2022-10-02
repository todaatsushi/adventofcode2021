package grid

import (
	"adventofcode2021/aoc11/src/utils"
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

func (g *Grid) Step() int {
	flashes := 0
	toVisit := utils.NewQueue()
	flashed := make(map[[2]int]bool)

	var newCoords [2]int

	for {
		current, err := toVisit.Next()
		if err != nil {
			break
		}
		newCoords = [2]int{current[0], current[1]}

		if g.Board[newCoords[0]][newCoords[1]] == 9 {
			flashed[newCoords] = true
			g.Board[newCoords[0]][newCoords[1]] = 0
			flashes += 1

			for _, move := range g.getMoves(newCoords) {
				toVisit.Add([2]int{move[0], move[1]})
			}
		} else {
			if _, ok := flashed[newCoords]; ok {
				// Max one flash a turn
				continue
			}
			g.Board[newCoords[0]][newCoords[1]] += 1
		}
	}
	g.Flashes += flashes
	return flashes
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
