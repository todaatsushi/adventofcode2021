package readings

import "testing"

func TestGetRawGammaRate(t *testing.T) {
	allReadings := []Reading{
		Reading{value: []int64{0, 1, 1, 0, 0}},
		Reading{value: []int64{1, 0, 1, 0, 0}},
		Reading{value: []int64{0, 1, 1, 0, 0}},
		Reading{value: []int64{1, 0, 1, 1, 0}},
		Reading{value: []int64{0, 1, 1, 0, 1}},
	}

	expectedGammaRateRaw := "11100"
	gammaRateRaw, _ := getRawGammaAndEpsilon(allReadings)

	if expectedGammaRateRaw != gammaRateRaw {
		t.Fatal("Different gamma rates calculated:", expectedGammaRateRaw, gammaRateRaw)
	}

}

func TestGetRawEpsilonRate(t *testing.T) {
	allReadings := []Reading{
		Reading{value: []int64{0, 1, 1, 0, 0}},
		Reading{value: []int64{1, 0, 1, 0, 0}},
		Reading{value: []int64{0, 1, 1, 0, 0}},
		Reading{value: []int64{1, 0, 1, 1, 0}},
		Reading{value: []int64{0, 1, 1, 0, 1}},
	}

	expectedEpsilonRateRaw := "00011"
	_, epsilonRateRaw := getRawGammaAndEpsilon(allReadings)

	if expectedEpsilonRateRaw != epsilonRateRaw {
		t.Fatal("Different epsilon rate rates calculated:", expectedEpsilonRateRaw, epsilonRateRaw)
	}

}
