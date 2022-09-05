package main

import (
	"adventofcode2021/aoc8/part1"
	"adventofcode2021/aoc8/part2"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	part1Answer := part1.Solve(filename)

	fmt.Println("part 1:", part1Answer)

	part2Answer := part2.Solve(filename)
	fmt.Println("part 2:", part2Answer)
}
