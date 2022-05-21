package readings

import "log"

type Tally struct {
	zero int64
	one  int64
}

func (tally *Tally) reset() {
	tally.one = 0
	tally.zero = 0
}

func (tally *Tally) getMax() int64 {
	if tally.one == tally.zero {
		log.Fatal("Can't handle same tally")
	}
	if tally.one < tally.zero {
		return 1
	}
	return 0
}
