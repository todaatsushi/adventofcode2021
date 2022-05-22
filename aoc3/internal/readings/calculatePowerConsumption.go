package readings

import (
	"log"
	"strconv"
)

func getRawGammaAndEpsilon(allReadings []Reading) ([]int64, []int64) {
	gammaRateRaw := make([]int64, 0)
	epsilonRateRaw := make([]int64, 0)

	tally := Tally{one: 0, zero: 0}
	bits := len(allReadings[0].value)

	var max int64
	var min int64
	var err error

	for i := 0; i < bits; i++ {
		tally.reset()
		tally.tallyReadingsAtPosition(allReadings, i)
		max, err = tally.getMax()
		min, err = tally.getMin()

		if err != nil {
			log.Fatal("Couldn't get max or min.")
		}

		gammaRateRaw = append(gammaRateRaw, max)
		epsilonRateRaw = append(epsilonRateRaw, min)
	}
	return gammaRateRaw, epsilonRateRaw
}

func GetGammaAndEpsilon(allReadings []Reading) (int64, int64) {
	gammaRateRaw, epsilonRateRaw := getRawGammaAndEpsilon(allReadings)
	gammaRateString := fmtReadingValue(gammaRateRaw)
	epsilonRateString := fmtReadingValue(epsilonRateRaw)

	gammaRate, err := strconv.ParseInt(gammaRateString, 2, 64)
	if err != nil {
		log.Fatal("Couldn't convert gammaRateRaw", gammaRateString)
	}
	epsilonRate, err := strconv.ParseInt(epsilonRateString, 2, 64)
	if err != nil {
		log.Fatal("Couldn't convert epsilonRateRaw", epsilonRateString)
	}
	return gammaRate, epsilonRate
}
