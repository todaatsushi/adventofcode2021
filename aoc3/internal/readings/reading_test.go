package readings

import (
	"testing"
)

func TestReadingPush(t *testing.T) {
	reading := Reading{}
	if len(reading.value) != 0 {
		t.Fatalf("Reading isn't empty")
	}

	reading.push(1)

	if len(reading.value) != 1 {
		t.Fatalf("Reading value len should be 1")
	}

	if reading.value[0] != 1 {
		t.Fatalf("Value should be one")
	}
}

func TestReadingPop(t *testing.T) {
	reading := Reading{value: []int64{1, 0}}

	if len(reading.value) != 2 {
		t.Fatal("Len of reading.value is not 2, after initialising on 2. Currently ", len(reading.value))
	}

	var first int64
	first = reading.pop()

	if len(reading.value) != 1 {
		t.Fatal("Len of reading.value is not 1, after popping on 2. Currently ", len(reading.value))
	}
	if first != 1 {
		t.Fatal("Popped value should be 1, got ", first)
	}
}
