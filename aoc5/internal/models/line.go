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

func (l *Line) getDiagonalPoints() []*Coordinate {
	start, end := l.getStartAndEndPoints()

	coordinates := make([]*Coordinate, end.x-start.x)
	coordinates[0] = start
	coordinates[len(coordinates)-1] = end

	inc := 1
	if end.y < start.y {
		inc = -1
	}

	x := start.x + 1
	y := start.y + inc

	for i := 1; i < end.x-start.x; i++ {
		coordinates[i] = &Coordinate{x: x, y: y}
		x += 1
		y += inc
	}
	return coordinates
}
