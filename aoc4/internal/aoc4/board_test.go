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
	var current int
	var expectedVal int
	total := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			current = board.nums[r][c]
			expectedVal = expected[r][c]

			if current != expectedVal {
				t.Fatal("Values aren't the same:", current, expected[r][c])
			}
			total += current
		}
	}

	if board.complete == true {
		t.Fatal("Board is marked as complete when it shouldn't")
	}
	if board.score != total {
		t.Fatal("Board total is incorrect.")
	}
}

func TestCheckBoard(t *testing.T) {
	input := `22 13 17 11  0
	 8  2 23  4 24
	21  9 14 16  7
	 6 10  3 18  5
	 1 12 20 15 19`

	board := newBoard(input)
	board.check(22)

	if board.score != (300 - 22) {
		t.Fatal("Scores don't match up: ", board.score)
	}
	if board.complete == true {
		t.Fatal("Shouldn't be marked as complete")
	}
	if board.rows[0] != 1 {
		t.Fatal("Row count not incremented")
	}
	if board.cols[0] != 1 {
		t.Fatal("Col count not incremented")
	}
}

func TestCheckBoardCompleteRow(t *testing.T) {
	input := `0 0 0 0 0
	 1 1 1 1 1
	 1 1 1 1 1
	 1 1 1 1 1
	 1 1 1 1 1`

	board := newBoard(input)
	board.check(0)

	if board.score != 20 {
		t.Fatal("Scores don't match up: ", board.score)
	}
	if board.complete == false {
		t.Fatal("Should be marked as complete")
	}
	if board.rows[0] != 5 {
		t.Fatal("Row count not completed")
	}
	if board.cols[0] != 1 {
		t.Fatal("Col count not incremented")
	}
}

func TestCheckBoardCompleteCol(t *testing.T) {
	input := `0 1 1 1 1
	 0 1 1 1 1
	 0 1 1 1 1
	 0 1 1 1 1
	 0 1 1 1 1`

	board := newBoard(input)
	board.check(0)

	if board.score != 20 {
		t.Fatal("Scores don't match up: ", board.score)
	}
	if board.complete == false {
		t.Fatal("Should be marked as complete")
	}
	if board.rows[0] != 1 {
		t.Fatal("Row count not completed")
	}
	if board.cols[0] != 5 {
		t.Fatal("Col count not incremented")
	}
}
