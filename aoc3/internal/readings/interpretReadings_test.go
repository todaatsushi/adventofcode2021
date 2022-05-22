package readings

import "testing"

func TestGetRawGammaRate(t *testing.T) {
	allReadings := []Reading{
		Reading{value: []int64{0, 1, 1, 0, 1}},
		Reading{value: []int64{1, 0, 1, 0, 0}},
		Reading{value: []int64{0, 1, 1, 0, 0}},
	}

	expectedGammaRateRaw := []int64{0, 1, 1, 0, 0}
	gammaRateRaw, _ := getRawGammaAndEpsilon(allReadings)

	for i := 0; i < 5; i++ {
		if expectedGammaRateRaw[i] != gammaRateRaw[i] {
			t.Fatal("Different gamma rates calculated:", expectedGammaRateRaw, gammaRateRaw)
		}
	}
}

func TestGetRawEpsilonRate(t *testing.T) {
	allReadings := []Reading{
		Reading{value: []int64{0, 1, 1, 0, 0}},
		Reading{value: []int64{1, 0, 1, 0, 0}},
		Reading{value: []int64{0, 1, 1, 0, 0}},
	}

	expectedEpsilonRateRaw := []int64{1, 0, 0, 1, 1}
	_, epsilonRateRaw := getRawGammaAndEpsilon(allReadings)
	for i := 0; i < 5; i++ {
		if expectedEpsilonRateRaw[i] != epsilonRateRaw[i] {
			t.Fatal("Different epsilon rates calculated:", expectedEpsilonRateRaw, epsilonRateRaw)
		}
	}
}
