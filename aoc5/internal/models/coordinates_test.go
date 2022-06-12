package models

import (
	"testing"
)

func TestReadCoordinates(t *testing.T) {
	input := "0,9 -> 5,9"
	start, end := readCoordinates(input)

	if start.x != 0 {
		t.Fatal("Start coordinate x value incorrect. Wanted 0, got ", start.x)
	}
	if start.y != 9 {
		t.Fatal("Start coordinate y value incorrect. Wanted 9, got ", start.y)
	}

	if end.x != 5 {
		t.Fatal("End coordinate x value incorrect. Wanted 5, got ", end.x)
	}
	if end.y != 9 {
		t.Fatal("Start coordinate y value incorrect. Wanted 9, got ", end.y)
	}

}
