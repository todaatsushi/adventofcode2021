package readings

func getRawGammaAndEpsilon(allReadings []Reading) ([]int64, []int64) {
	gammaRateRaw := make([]int64, 0)
	epsilonRateRaw := make([]int64, 0)

	tally := Tally{one: 0, zero: 0}
	var bit int64
	bits := len(allReadings[0].value)

	for i := 0; i < bits; i++ {
		tally.reset()
		for _, reading := range allReadings {
			bit = reading.value[i]
			if bit == 1 {
				tally.one = tally.one + 1
			} else {
				tally.zero = tally.zero + 1
			}
		}
		gammaRateRaw = append(gammaRateRaw, tally.getMax())
		epsilonRateRaw = append(epsilonRateRaw, tally.getMin())
	}
	return gammaRateRaw, epsilonRateRaw
}
