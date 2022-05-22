package readings

import (
	"testing"
)

func TestGetMostCommonValueAtPosition(t *testing.T) {
	allReadings := createTestReadings()
	expected := []int64{0, 1, 1}
	var mostCommonValueAtPostion int64

	for i := 0; i < 3; i++ {
		mostCommonValueAtPostion = getMostCommonValueAtPosition(allReadings, i)
		if mostCommonValueAtPostion != expected[i] {
			t.Fatal("Values don't match.", mostCommonValueAtPostion, expected[i])
		}
	}
}
func TestGetLeastCommonValueAtPosition(t *testing.T) {
	allReadings := createTestReadings()
	expected := []int64{1, 0, 0}
	var leastCommonValueAtPostion int64

	for i := 0; i < 3; i++ {
		leastCommonValueAtPostion = getLeastCommonValueAtPosition(allReadings, i)
		if leastCommonValueAtPostion != expected[i] {
			t.Fatal("Values don't match.", leastCommonValueAtPostion, expected[i])
		}
	}
}

func TestGetO2Reading(t *testing.T) {
	allReadings := createTestReadings()
	expected := int64(13)
	O2Reading := GetO2Reading(allReadings, 0)
	if O2Reading != expected {
		t.Fatal("O2 readings didn't match expected:", O2Reading, expected)
	}
}

func TestGetCO2Readings(t *testing.T) {
	allReadings := createTestReadings()
	expected := int64(20)
	CO2Reading := GetCO2Reading(allReadings, 0)

	if CO2Reading != expected {
		t.Fatal("CO2 readings didn't match expected:", CO2Reading, expected)
	}

}
