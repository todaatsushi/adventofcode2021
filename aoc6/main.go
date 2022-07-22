package main

import (
	"adventofcode2021/aoc6/src"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	fishAges := src.ReadInput(filename)
	fmt.Println(fishAges)
}
