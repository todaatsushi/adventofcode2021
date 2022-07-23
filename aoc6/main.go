package main

import (
	"adventofcode2021/aoc6/src"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	daysStr := os.Args[2]
	days, _ := strconv.Atoi(daysStr)

	timeline := src.SpawnFishFromInput(filename)
	src.SimulateFishGrowth(timeline, days)
}
