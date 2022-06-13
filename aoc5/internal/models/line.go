package models

import (
	"adventofcode2021/aoc5/internal/utils"
	"sync"
)

type Line struct {
	start      *Coordinate
	end        *Coordinate
	isVertical bool
}

func newLine(start *Coordinate, end *Coordinate) *Line {
	isVertical := start.x == end.x

	return &Line{start: start, end: end, isVertical: isVertical}
}

func ReadLines(input []string) []*Line {
	lines := make([]*Line, len(input))
	wg := sync.WaitGroup{}

	for i, coords := range input {
		wg.Add(1)
		go func(lines []*Line, i int, coords string) {
			defer wg.Done()
			start, end := readCoordinates(coords)
			lines[i] = newLine(start, end)
		}(lines, i, coords)
	}

	wg.Wait()
	return lines
}

func (l *Line) getMaxX() int {
	return utils.Max(l.start.x, l.end.x)
}

func (l *Line) getMaxY() int {
	return utils.Max(l.start.y, l.end.y)
}
