package aoc5

import "testing"

func TestReadLines(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
	}
	lines := readLines(input)

	if len(lines) != 4 {
		t.Fatal("Unexpected number of lines:", len(lines))
	}
}
