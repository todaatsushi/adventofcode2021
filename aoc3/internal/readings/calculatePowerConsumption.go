package readings

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func fmtRate(rateSlice []int64) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(rateSlice)), ""), "[]")
}

func getRawGammaAndEpsilon(allReadings []Reading) ([]int64, []int64) {
	gammaRateRaw := make([]int64, 0)
	epsilonRateRaw := make([]int64, 0)

	tally := Tally{one: 0, zero: 0}
	var bit int64
	bits := len(allReadings[0].value)

	var max int64
	var min int64
	var err error

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
	gammaRateString := fmtRate(gammaRateRaw)
	epsilonRateString := fmtRate(epsilonRateRaw)

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
