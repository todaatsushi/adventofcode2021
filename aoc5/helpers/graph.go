package helpers

import (
	"fmt"
)

type Graph struct {
	Points [][]int
}

func (g *Graph) print(pretty bool) {
	for r := 0; r < len((*g).Points); r++ {
		for c := 0; c < len((*g).Points[r]); c++ {
			if pretty == true {
				fmt.Printf("%d ", (*g).Points[c][r])
			} else {
				fmt.Printf("%d ", (*g).Points[r][c])
			}
		}
		fmt.Print("\n")
	}
}

func NewGraph(lines []Line) *Graph {
	x := 0
	y := 0
	tempX := 0
	tempY := 0

	for _, l := range lines {
		tempX = max(l.StartX, l.EndX)
		tempY = max(l.StartY, l.EndY)

		x = max(x, tempX)
		y = max(y, tempY)
	}
	x++
	y++

	points := make([][]int, x)
	for i := 0; i < x; i++ {
		points[i] = make([]int, y)
	}
	graph := Graph{Points: points}
	return &graph
}
