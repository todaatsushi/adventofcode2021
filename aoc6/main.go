package main

import (
	"adventofcode2021/aoc6/src"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	daysStr := os.Args[2]
	days, _ := strconv.Atoi(daysStr)

	fishAges := src.GetFishAgesFromInput(filename)
	fish := src.FishFromAges(fishAges)

	total := src.GetNumFishAfterDays(fish, days)
	fmt.Println("Num of fish:", total)
}
