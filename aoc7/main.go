package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func parseCrabsFromFile(filename string) []int {
	content, _ := os.ReadFile(filename)
	rawContent := string(content)
	positions := strings.Split(rawContent, ",")

	crabs := make([]int, len(positions))
	wg := sync.WaitGroup{}
	for i, p := range positions {
		wg.Add(1)
		p = strings.TrimSpace(p)
		go func(crabs *[]int, index int, val string) {
			defer wg.Done()
			asNum, _ := strconv.Atoi(val)
			(*crabs)[index] = asNum
		}(&crabs, i, p)
	}
	wg.Wait()

	sort.Ints(crabs)
	return crabs
}

func createLocationMap(crabs []int) map[int]int {
	lm := make(map[int]int)
	for _, c := range crabs {
		lm[c]++
	}
	return lm
}

func absDiffInt(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func findOptimalConversionPoint(locationMap *map[int]int, start int, end int) int {
	smallestMoves := math.MaxInt

	for point := start; point < end+1; point++ {
		total := 0
		for location, crabCount := range *locationMap {
			if location == point {
				continue
			}
			distance := absDiffInt(point, location)
			total += distance * crabCount
		}

		if smallestMoves > total {
			smallestMoves = total
		}
	}

	return smallestMoves
}

func main() {
	filename := os.Args[1]
	crabs := parseCrabsFromFile(filename)
	locationMap := createLocationMap(crabs)

	bestPlacement := findOptimalConversionPoint(&locationMap, crabs[0], crabs[len(crabs)-1])
	fmt.Println("Min moves:", bestPlacement)
}
