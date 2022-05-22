package readings

import "testing"

func createTestReadings() []Reading {
	return []Reading{
		Reading{value: []int64{0, 1, 1, 0, 1}},
		Reading{value: []int64{1, 0, 1, 0, 0}},
		Reading{value: []int64{0, 1, 1, 0, 0}},
	}
}

func TestGetRawGammaRate(t *testing.T) {
	allReadings := createTestReadings()
	expectedGammaRateRaw := []int64{0, 1, 1, 0, 0}
	gammaRateRaw, _ := getRawGammaAndEpsilon(allReadings)

	for i := 0; i < 5; i++ {
		if expectedGammaRateRaw[i] != gammaRateRaw[i] {
			t.Fatal("Different gamma rates calculated:", expectedGammaRateRaw, gammaRateRaw)
		}
	}
}

func TestGetRawEpsilonRate(t *testing.T) {
	allReadings := createTestReadings()
	expectedEpsilonRateRaw := []int64{1, 0, 0, 1, 1}
	_, epsilonRateRaw := getRawGammaAndEpsilon(allReadings)
	for i := 0; i < 5; i++ {
		if expectedEpsilonRateRaw[i] != epsilonRateRaw[i] {
			t.Fatal("Different epsilon rates calculated:", expectedEpsilonRateRaw, epsilonRateRaw)
		}
	}
}

func TestFmtRate(t *testing.T) {
	rateReading := []int64{0, 1, 1, 0, 0}
	rateReadingString := fmtRate(rateReading)
	expectedRateReadingString := "01100"

	if rateReadingString != expectedRateReadingString {
		t.Fatal("Rate reading not pasted correctly into a string expected '01100', got:", rateReadingString)
	}
}
