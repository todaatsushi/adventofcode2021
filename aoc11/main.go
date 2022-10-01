package main

import (
	"adventofcode2021/aoc11/src/utils"
	"fmt"
)

func main() {
	filepath := "inputs/test.txt"
	grid := utils.GetGridFromFile(filepath)

	fmt.Println(grid)
}
