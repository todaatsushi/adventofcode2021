package src

import (
	"os"
	"strconv"
	"strings"
)

func GetFishAgesFromInput(filename string) []int {
	content, _ := os.ReadFile(filename)
	rawInput := string(content)
	rawInput = strings.TrimSpace(rawInput)

	fishStrings := strings.Split(rawInput, ",")
	fish := make([]int, len(fishStrings))

	var days int
	for i, fs := range fishStrings {
		fs = strings.TrimSpace(fs)
		days, _ = strconv.Atoi(fs)

		fish[i] = days
	}
	return fish
}
