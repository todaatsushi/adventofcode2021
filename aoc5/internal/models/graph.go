package models

import (
	"adventofcode2021/aoc5/internal/utils"
)

func getGraphSize(lines []*Line) (int, int) {
	x := 0
	y := 0

	for _, l := range lines {
		x = utils.Max(x, l.getMaxX())
		y = utils.Max(y, l.getMaxY())
	}
	return x + 1, y + 1
}

func CreateGraph(lines []*Line) *[][]int {
	xSize, ySize := getGraphSize(lines)
	graph := make([][]int, xSize)

	for i := 0; i < xSize; i++ {
		graph[i] = make([]int, ySize)
	}
	return &graph
}

func drawLine(graph *[][]int, line *Line) {
	if line.isVertical {
		for i := line.start.y; i <= line.end.y; i++ {
			(*graph)[line.start.x][i] += 1
		}
	} else {
		for i := line.start.x; i <= line.end.x; i++ {
			(*graph)[i][line.start.y] += 1
		}
	}
}

func ReadLinesToGraph(graph *[][]int, lines []*Line) {
	for _, l := range lines {
		drawLine(graph, l)
	}
}

func GetPointsWithValOverX(graph *[][]int, x int) int {
	count := 0
	for r := 0; r < 10; r++ {
		for c := 0; c < 10; c++ {
			if (*graph)[r][c] >= x {
				count += 1
			}
		}
	}
	return count
}
