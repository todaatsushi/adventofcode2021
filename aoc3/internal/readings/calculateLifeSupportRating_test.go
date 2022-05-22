package readings

import (
	"strconv"
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

func TestGetO2Readings(t *testing.T) {
	allReadings := createTestReadings()
	expected := int64(13)
	O2Reading := GetO2Readings(allReadings, 0)
	O2ReadingStr := fmtRate(O2Reading.value)
	intO2Reading, _ := strconv.ParseInt(O2ReadingStr, 2, 64)

	if intO2Reading != expected {
		t.Fatal("O2 readings didn't match expected:", intO2Reading, expected)
	}
}
