package part2

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"adventofcode2021/aoc8/utils"
)

func generateCharMap(input []string) map[string]int {
	charMap := utils.NewDict()

	// ID unique patterns
	for _, pattern := range input {
		pattern = utils.SortString(pattern)
		switch len(pattern) {
		case 2:
			charMap.MapVal(pattern, 1)
		case 4:
			charMap.MapVal(pattern, 4)
		case 3:
			charMap.MapVal(pattern, 7)
		case 7:
			charMap.MapVal(pattern, 8)
		}
	}

	for _, pattern := range input {
		pattern = utils.SortString(pattern)
		switch len(pattern) {
		case 5:
			// ID count of 5 (2, 3, 5)
			matchesWith1 := utils.NumCharsMatches(charMap.Index[1], pattern)
			if matchesWith1 == 2 {
				charMap.MapVal(pattern, 3)
			} else {
				matchesWith4 := utils.NumCharsMatches(charMap.Index[4], pattern)
				if matchesWith4 == 3 {
					charMap.MapVal(pattern, 5)
				} else {
					charMap.MapVal(pattern, 2)
				}
			}

		case 6:
			// ID count of 6 (9, 0, 6)
			// Always done after len 5s
			matchesWith4 := utils.NumCharsMatches(charMap.Index[4], pattern)
			matchesWith1 := utils.NumCharsMatches(charMap.Index[1], pattern)
			if matchesWith4 == 3 && matchesWith1 == 2 {
				charMap.MapVal(pattern, 0)
			} else {
				matchesWith1 := utils.NumCharsMatches(charMap.Index[1], pattern)
				if matchesWith1 == 2 {
					charMap.MapVal(pattern, 9)
				} else {
					charMap.MapVal(pattern, 6)
				}
			}
		}
	}

	return charMap.Table
}

func createOutputVal(charMap *map[string]int, output []string) int {
	var number string
	for _, pattern := range output {
		pattern = utils.SortString(pattern)
		number = fmt.Sprintf("%v%v", number, (*charMap)[pattern])
	}

	val, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal("Couldn't convert:", number)
	}
	return val
}

func Solve(filename string) int {
	content := utils.ReadInput(filename)

	var input []string
	var output []string

	var val int
	total := 0
	for _, line := range content {
		input = line[0]
		output = line[1]

		sort.Sort(utils.ByLength(input))
		charMap := generateCharMap(input)

		val = createOutputVal(&charMap, output)
		total += val
	}

	return total
}
