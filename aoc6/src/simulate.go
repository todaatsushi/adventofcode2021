package src

import (
	"os"
	"strconv"
	"strings"
)

func getFishAges(filename string) []int {
	content, _ := os.ReadFile(filename)
	rawInput := string(content)
	fishStrs := strings.Split(rawInput, ",")

	numFish := len(fishStrs)
	fish := make([]int, numFish)
	var f int

	for i, fs := range fishStrs {
		fs = strings.TrimSpace(fs)
		f, _ = strconv.Atoi(fs)
		fish[i] = f
	}
	return fish
}

func SpawnFishFromInput(filename string) []int {
	fish := getFishAges(filename)
	timeline := make([]int, 9)

	for _, age := range fish {
		timeline[age] += 1
	}
	return timeline
}
