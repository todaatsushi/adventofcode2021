package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Graph struct {
	points [][]int
}

func (g *Graph) print(pretty bool) {
	for r := 0; r < len((*g).points); r++ {
		for c := 0; c < len((*g).points[r]); c++ {
			if pretty == true {
				fmt.Printf("%d ", (*g).points[c][r])
			} else {
				fmt.Printf("%d ", (*g).points[r][c])
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
		tempX = max(l.startX, l.endX)
		tempY = max(l.startY, l.endY)

		x = max(x, tempX)
		y = max(y, tempY)
	}
	x++
	y++

	points := make([][]int, x)
	for i := 0; i < x; i++ {
		points[i] = make([]int, y)
	}
	graph := Graph{points: points}
	return &graph
}

type Line struct {
	startX int
	startY int
	endX   int
	endY   int
}

func (l *Line) isStraight() bool {
	return l.startX == l.endX || l.startY == l.endY
}

func getXAndY(lineStr string) (int, int) {
	coords := strings.Split(lineStr, ",")

	strX := strings.TrimSpace(coords[0])
	strY := strings.TrimSpace(coords[1])
	intX, _ := strconv.Atoi(strX)
	intY, _ := strconv.Atoi(strY)

	return intX, intY
}

func (l *Line) getStartAndEnd(axis string) (int, int) {
	if axis == "x" {
		if l.startX > l.endX {
			return l.endX, l.startX
		} else {
			return l.startX, l.endX
		}
	} else {
		if l.startY > l.endY {
			return l.endY, l.startY
		} else {
			return l.startY, l.endY
		}
	}
}

func newLine(inputs string) Line {
	parts := strings.Split(inputs, "->")
	startX, startY := getXAndY(parts[0])
	endX, endY := getXAndY(parts[1])
	return Line{startX: startX, startY: startY, endX: endX, endY: endY}
}

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Couldn't read file:", filename)
	}
	rawInput := string(content)
	rawInput = strings.TrimSpace(rawInput)
	input := strings.Split(rawInput, "\n")

	return input
}

func main() {
	filename := os.Args[1]

	input := readInput(filename)
	numLines := len(input)

	lines := make([]Line, numLines)
	for i, l := range input {
		lines[i] = newLine(l)
	}

	graph := *NewGraph(lines)
	tally := 0

	var start int
	var end int
	for _, l := range lines {
		if l.isStraight() == true {
			if l.startX == l.endX {
				start, end = l.getStartAndEnd("y")
				for i := start; i <= end; i++ {
					graph.points[l.startX][i]++
					if graph.points[l.startX][i] == 2 {
						tally++
					}
				}
			} else {
				start, end = l.getStartAndEnd("x")
				for i := start; i <= end; i++ {
					graph.points[i][l.startY]++
					if graph.points[i][l.startY] == 2 {
						tally++
					}
				}
			}
		}
	}

	graph.print(true)
	fmt.Println()
	fmt.Println(tally)
}
