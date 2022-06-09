package aoc4

import "testing"

func TestReadInput(t *testing.T) {
	input := `
	7,4,9,5,11

	14 21 17 24 4
	10 16 15 9 19
	18 8 23 26 20
	22 11 13 6 5
	2 0 12 3 7
	`

	calls, boards := ReadInput(input)

	expectedCalls := "7,4,9,5,11"
	expectedBoards := []string{`14 21 17 24 4
	10 16 15 9 19
	18 8 23 26 20
	22 11 13 6 5
	2 0 12 3 7`}

	if expectedCalls != calls {
		t.Fatal("Calls are incorrectly split")
	}
	if expectedBoards[0] != boards[0] {
		t.Fatal("Board is incorrectly split")
	}
}
