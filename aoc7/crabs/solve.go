package crabs

import "math"

func absDiffInt(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func getFuelCost(point int, location int) int {
	return absDiffInt(point, location)
}

func FindOptimalConversionPoint(locationMap *map[int]int, start int, end int) int {
	smallestMoves := math.MaxInt

	for point := start; point < end+1; point++ {
		total := 0
		for location, crabCount := range *locationMap {
			if location == point {
				continue
			}
			distance := getFuelCost(point, location)
			total += distance * crabCount
		}

		if smallestMoves > total {
			smallestMoves = total
		}
	}

	return smallestMoves
}
