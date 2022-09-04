package part1

import (
	"adventofcode2021/aoc8/utils"
)

func getCharCountMap() map[int]int {
	charCounts := map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}
	return charCounts
}

func Solve(filename string) int {
	outputs := utils.ReadInput(filename)
	charCounts := getCharCountMap()

	tally := map[int]int{1: 0, 4: 0, 7: 0, 8: 0}

	var outputSet []string
	for _, io := range outputs {
		outputSet = io[1]
		for _, output := range outputSet {
			if num, ok := charCounts[len(output)]; ok {
				tally[num]++
			}
		}
	}
	total := 0
	for _, count := range tally {
		total += count
	}
	return total
}
