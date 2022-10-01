package utils

import (
	"log"
	"os"
	"strconv"
	"strings"

	"adventofcode2021/aoc11/src/grid"
)

func readFile(filepath string) string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Couldn't read file! Path: ", filepath)
	}

	contentStr := string(content)
	contentStr = strings.TrimSpace(contentStr)
	return contentStr
}

func GetGridFromFile(filepath string) grid.Grid {
	contentStr := readFile(filepath)
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

	grid := grid.NewGrid(board)
	return grid
}
