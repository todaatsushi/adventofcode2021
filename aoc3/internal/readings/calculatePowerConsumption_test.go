package readings

import (
	"testing"
)

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

func TestFmtReadingValue(t *testing.T) {
	valueReading := []int64{0, 1, 1, 0, 0}
	valueReadingString := fmtReadingValue(valueReading)
	expectedRateReadingString := "01100"

	if valueReadingString != expectedRateReadingString {
		t.Fatal("Rate reading not pasted correctly into a string expected '01100', got:", valueReadingString)
	}
}

func TestGetGammaRate(t *testing.T) {
	allReadings := createTestReadings()
	gammaRate, _ := GetGammaAndEpsilon(allReadings)
	expectedGammaRate := 12

	if gammaRate != int64(expectedGammaRate) {
		t.Fatal("Gamme rate is not as expected. Expected 22, got ", gammaRate)
	}
}

func TestGetEpsilonRate(t *testing.T) {
	allReadings := createTestReadings()
	_, epsilonRate := GetGammaAndEpsilon(allReadings)
	expectedEpsilonRate := 19

	if epsilonRate != int64(expectedEpsilonRate) {
		t.Fatal("Epsilon rate is not as expected. Expected 9, got ", epsilonRate)
	}

}
