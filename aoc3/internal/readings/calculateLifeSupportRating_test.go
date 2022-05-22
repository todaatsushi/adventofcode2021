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
