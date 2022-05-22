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
	max := tally.getMax()
	if max != 0 {
		log.Fatal("Didn't get largest of the tallies, expected 0, got", max)
	}
}

func TestTallyGetMin(t *testing.T) {
	tally := Tally{one: 1, zero: 2}
	min := tally.getMin()
	if min != 1 {
		log.Fatal("Didn't get smallest of the tallies, expected 0, got", min)
	}
}
