package aoc5

import (
	"adventofcode2021/aoc5/internal/models"
	"fmt"
)

func Solve(filename string, val int, withDiagonals bool) {
	input := readInput(filename)
	lines := models.ReadLines(input)
	graph := models.NewGraph(lines)

	graph.ReadLines(withDiagonals)
	tally := graph.PointsWithOverXNumberOfLines(val)

	fmt.Printf("Number of points in graph over %d is %d\n", val, tally)
	graph.Describe(false)
}
