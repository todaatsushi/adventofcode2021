package src

type Fish struct {
	internalTimer int
}

func createFish(internalTimer int) *Fish {
	return &Fish{internalTimer: internalTimer}
}
