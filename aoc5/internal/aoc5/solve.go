package aoc5

import (
	"adventofcode2021/aoc5/internal/models"
	"fmt"
)

func Solve(filename string, val int) {
	input := readInput(filename)
	lines := models.ReadLines(input)
	graph := models.CreateGraph(lines)

	models.ReadLinesToGraph(graph, lines)
	tally := models.GetPointsWithValOverX(graph, val)

	fmt.Printf("Number of points in graph over %d is %d\n", val, tally)
}
