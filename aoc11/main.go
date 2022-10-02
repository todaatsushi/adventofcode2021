package main

import (
	"adventofcode2021/aoc11/src/grid"
	"adventofcode2021/aoc11/src/utils"
	"fmt"
)

func main() {
	filepath, numSteps, display := utils.GetArgs()

	grid := grid.GetGridFromFile(filepath)
	grid.FlashXSteps(numSteps, display)
	if display == false {
		fmt.Println("Flashes:", grid.Flashes)
	}
}
