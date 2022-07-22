package helpers

import (
	"strconv"
	"strings"
)

type Line struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
}

func (l *Line) IsStraight() bool {
	return l.StartX == l.EndX || l.StartY == l.EndY
}

func getXAndY(lineStr string) (int, int) {
	coords := strings.Split(lineStr, ",")

	strX := strings.TrimSpace(coords[0])
	strY := strings.TrimSpace(coords[1])
	intX, _ := strconv.Atoi(strX)
	intY, _ := strconv.Atoi(strY)

	return intX, intY
}

func (l *Line) GetStartAndEnd(axis string) (int, int) {
	if axis == "x" {
		if l.StartX > l.EndX {
			return l.EndX, l.StartX
		} else {
			return l.StartX, l.EndX
		}
	} else {
		if l.StartY > l.EndY {
			return l.EndY, l.StartY
		} else {
			return l.StartY, l.EndY
		}
	}
}

func newLine(inputs string) Line {
	parts := strings.Split(inputs, "->")
	startX, startY := getXAndY(parts[0])
	endX, endY := getXAndY(parts[1])
	return Line{StartX: startX, StartY: startY, EndX: endX, EndY: endY}
}

func ReadLines(rawLines []string) []Line {
	numLines := len(rawLines)

	lines := make([]Line, numLines)
	for i, l := range rawLines {
		lines[i] = newLine(l)
	}
	return lines
}
