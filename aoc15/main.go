package main

import (
	"fmt"
	"log"
	"os"

	"adventofcode2021/aoc15/src/grid"
	"adventofcode2021/aoc15/src/input"
)

func main() {
	numArgs := len(os.Args)
	if numArgs < 2 {
		log.Fatal("Not enough args.")
	}

	filepath := os.Args[1]
	content := input.ReadInput(filepath)

	fmt.Println(content)

	g := grid.NewGrid(content)
	fmt.Println(g)
}
