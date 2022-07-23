package crabs

import "math"

func absDiffInt(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func FindOptimalConversionPoint(locationMap *map[int]int, start int, end int) int {
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
