package main

import (
	"adventofcode2021/aoc7/crabs"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	c := crabs.ParseCrabsFromFile(filename)
	locationMap := crabs.CreateLocationMap(c)
	bestPlacement := crabs.FindOptimalConversionPoint(&locationMap, c[0], c[len(c)-1])
	fmt.Println("Min moves:", bestPlacement)
}
