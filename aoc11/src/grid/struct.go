package grid

import (
	"fmt"
)

type Grid struct {
	Board               [10][10]int
	Flashes             int
	coordinateModifiers [8][2]int
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
