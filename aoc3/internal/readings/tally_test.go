package readings

import (
	"log"
	"testing"
)

func TestTallyReset(t *testing.T) {
	tally := Tally{one: 1, zero: 2}

	tally.reset()
	if tally.one != 0 {
		log.Fatal("Reset didn't reset one, got", tally.one)
	}
	if tally.zero != 0 {
		log.Fatal("Reset didn't reset zero, got", tally.zero)
	}
}

func TestTallyGetMax(t *testing.T) {
	tally := Tally{one: 1, zero: 2}
	max, _ := tally.getMax()
	if max != 0 {
		log.Fatal("Didn't get largest of the tallies, expected 0, got", max)
	}

	tally.one = 2
	max, err := tally.getMax()
	if err == nil || max != -1 {
		t.Fatal("Error not raised with equal tallies")
	}
}

func TestTallyGetMin(t *testing.T) {
	tally := Tally{one: 1, zero: 2}
	min, _ := tally.getMin()
	if min != 1 {
		t.Fatal("Didn't get smallest of the tallies, expected 0, got", min)
	}

	tally.one = 2
	min, err := tally.getMin()
	if err == nil || min != -1 {
		t.Fatal("Error not raised with equal tallies")
	}

}
