package grid

import (
	"adventofcode2021/aoc11/src/utils"
	"log"
	"strconv"
	"strings"
)

func newGrid(board [10][10]int) Grid {
	grid := Grid{Board: board, Flashes: 0, Steps: 0, coordinateModifiers: utils.GetMoves()}
	return grid
}

func GetGridFromFile(filepath string) Grid {
	contentStr := utils.ReadFile(filepath)
	lines := strings.Split(contentStr, "\n")

	var board [10][10]int
	for row, line := range lines {
		octopuses := strings.Split(line, "")
		for col, octopus := range octopuses {
			val, err := strconv.Atoi(octopus)
			if err != nil {
				log.Fatal("Couldn't convert value to int:", octopus)
			}
			board[row][col] = val
		}
	}

	grid := newGrid(board)
	return grid
}
