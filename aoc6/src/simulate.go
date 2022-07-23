package src

import (
	"fmt"
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

func simulateDay(timeline *[]int) {
	var todayFishTotal int
	dayBeforeFishTotal := (*timeline)[0]

	for day := 8; day >= 0; day-- {
		todayFishTotal = (*timeline)[day]
		(*timeline)[day] = dayBeforeFishTotal
		dayBeforeFishTotal = todayFishTotal

		if day == 0 {
			(*timeline)[6] += todayFishTotal
		}
	}
}

func SimulateFishGrowth(timeline []int, days int) {
	elapsed := 0
	for {
		if elapsed == days {
			break
		}
		simulateDay(&timeline)
		elapsed++
	}

	total := 0
	for _, day := range timeline {
		total += day
	}
	fmt.Println("Num fish:", total)
}
