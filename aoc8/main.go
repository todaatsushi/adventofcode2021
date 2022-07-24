package main

import (
	"adventofcode2021/aoc8/part1"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	part1Answer := part1.Solve(filename)

	fmt.Println("part 1:", part1Answer)
}
