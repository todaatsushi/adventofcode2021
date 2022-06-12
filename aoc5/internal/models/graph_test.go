package models

import "testing"

func TestGetGraphSize(t *testing.T) {
	input := []string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
	}
	lines := ReadLines(input)

	x, y := getGraphSize(lines)
	if x != 9 {
		t.Fatal("Expected 9 for max X, got", x)
	}
	if y != 9 {
		t.Fatal("Expected 9 for max Y, got", y)
	}
}
