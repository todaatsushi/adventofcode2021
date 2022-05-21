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
