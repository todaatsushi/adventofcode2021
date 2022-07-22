package src

type Fish struct {
	internalTimer int
}

func createFish(internalTimer int) *Fish {
	return &Fish{internalTimer: internalTimer}
}

func FishFromAges(timers []int) []*Fish {
	fish := []*Fish{}

	var created *Fish
	for _, timer := range timers {
		created = createFish(timer)
		fish = append(fish, created)
	}
	return fish
}
