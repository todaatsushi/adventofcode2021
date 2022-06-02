package aoc4

import (
	"strconv"
	"strings"
)

type Number struct {
	value  int
	marked bool
}

type Board struct {
	nums [5][5]Number
}

func newBoard(input string) Board {
	rows := strings.Split(input, "\n")
	board := Board{}

	var vals []string
	var boardNum Number
	var colNum int
	for rowNum, row := range rows {
		row = strings.TrimSpace(row)
		vals = strings.Split(row, " ")
		colNum = 0

		for _, val := range vals {
			numVal, err := strconv.Atoi(val)
			if err == nil {
				boardNum = Number{value: numVal, marked: false}
				board.nums[rowNum][colNum] = boardNum
				colNum++
			}
		}
	}
	return board
}
