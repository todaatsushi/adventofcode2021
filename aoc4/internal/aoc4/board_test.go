package aoc4

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	input := `22 13 17 11  0
	 8  2 23  4 24
	21  9 14 16  7
	 6 10  3 18  5
	 1 12 20 15 19`
	expected := [][]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	}

	board := newBoard(input)
	var current Number
	var expectedVal int
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			current = board.nums[r][c]
			expectedVal = expected[r][c]

			if current.value != expectedVal {
				t.Fatal("Values aren't the same:", current, expected[r][c])
			}
			if current.marked == true {
				t.Fatalf("Marked should init as false")
			}
		}
	}
}
