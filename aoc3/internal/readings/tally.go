package readings

import (
	"errors"
)

type Tally struct {
	zero int64
	one  int64
}

func (tally *Tally) reset() {
	tally.one = 0
	tally.zero = 0
}

func (tally *Tally) getMax() (int64, error) {
	if tally.one == tally.zero {
		return -1, errors.New("Tally is equal")
	}
	if tally.one > tally.zero {
		return 1, nil
	}
	return 0, nil
}

func (tally *Tally) getMin() (int64, error) {
	if tally.one == tally.zero {
		return -1, errors.New("Tally is equal")
	}
	if tally.one < tally.zero {
		return 1, nil
	}
	return 0, nil

}
