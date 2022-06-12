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
	return x, y
}
