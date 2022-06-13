package models

import (
	"adventofcode2021/aoc5/internal/utils"
	"fmt"
)

type Graph struct {
	points [][]int
	lines  []*Line
}

func NewGraph(lines []*Line) *Graph {
	xSize, ySize := getGraphSize(lines)
	points := make([][]int, xSize)

	for i := 0; i < xSize; i++ {
		points[i] = make([]int, ySize)
	}
	graph := Graph{lines: lines, points: points}
	return &graph
}

func getGraphSize(lines []*Line) (int, int) {
	x := 0
	y := 0

	for _, l := range lines {
		x = utils.Max(x, l.getMaxX())
		y = utils.Max(y, l.getMaxY())
	}
	return x + 1, y + 1
}

func (g *Graph) getSize() (int, int) {
	return getGraphSize(g.lines)
}

func (g *Graph) drawLine(l *Line) {
	var startPoint *Coordinate
	var endPoint *Coordinate

	if l.isReversed {
		startPoint = l.end
		endPoint = l.start
	} else {
		startPoint = l.start
		endPoint = l.end
	}

	if l.isVertical {
		for i := startPoint.y; i <= endPoint.y; i++ {
			(*g).points[startPoint.x][i] += 1
		}
	} else {
		for i := startPoint.x; i <= endPoint.x; i++ {
			(*g).points[i][l.start.y] += 1
		}
	}

}

func (g *Graph) ReadLines() {
	for _, l := range g.lines {
		if l.isStraight() == true {
			g.drawLine(l)
		}
	}
}

func (g *Graph) Print() {
	for r := len((*g).points) - 1; r > 0; r-- {
		fmt.Println((*g).points[r])
	}

}

func (g *Graph) PointsWithOverXNumberOfLines(x int) int {
	count := 0
	for r := 0; r < len((*g).points); r++ {
		for c := 0; c < len((*g).points[r]); c++ {
			if (*g).points[r][c] >= x {
				count += 1
			}
		}
	}
	return count

}
