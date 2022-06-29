package models

import (
	"adventofcode2021/aoc5/internal/utils"
	"sync"
)

type Line struct {
	start *Coordinate
	end   *Coordinate

	isVertical bool
	isReversed bool
}

func newLine(start *Coordinate, end *Coordinate) *Line {
	isVertical := start.x == end.x
	isReversed := start.x > end.x || start.y > end.y

	return &Line{start: start, end: end, isVertical: isVertical, isReversed: isReversed}
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

func (l *Line) isStraight() bool {
	return l.start.x == l.end.x || l.start.y == l.end.y
}

func (l *Line) getStartAndEndPoints() (*Coordinate, *Coordinate) {
	lineIsStraight := l.isStraight()
	straightAndReversed := lineIsStraight == true && l.isReversed == true
	diagonalAndEndStartsFirst := lineIsStraight == false && l.start.x > l.end.x

	flipPoints := straightAndReversed || diagonalAndEndStartsFirst

	if flipPoints == true {
		return l.end, l.start
	}
	return l.start, l.end
}
