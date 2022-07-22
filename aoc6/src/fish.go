package src

type Fish struct {
	internalTimer int
}

func (f *Fish) age() bool {
	create := false
	if f.internalTimer == 0 {
		f.internalTimer = 6
		create = true
	} else {
		f.internalTimer--
	}
	return create
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

func GetNumFishAfterDays(fish []*Fish, days int) int {
	var create bool
	var newFish *Fish
	elapsed := 0

	for {
		if days == elapsed {
			break
		}
		for _, f := range fish {
			create = f.age()
			if create == true {
				newFish = createFish(8)
				fish = append(fish, newFish)
			}
		}
		elapsed++
		create = false
	}
	return len(fish)
}
