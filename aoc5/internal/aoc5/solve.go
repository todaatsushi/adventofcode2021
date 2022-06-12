package aoc5

import (
	"adventofcode2021/aoc5/internal/models"
	"fmt"
)

func Solve(filename string) {
	input := readInput(filename)
	lines := models.ReadLines(input)

	fmt.Println(lines, len(lines))
}
