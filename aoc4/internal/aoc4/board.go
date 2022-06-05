package aoc4

import (
	"strconv"
	"strings"
	"sync"
)

type Number struct {
	value  int
	marked bool
}

type Board struct {
	nums [5][5]Number

	rows [5]int
	cols [5]int
}

func createEmptyCounts() [5]int {
	return [5]int{0, 0, 0, 0, 0}
}

func newBoard(input string) Board {
	rows := strings.Split(input, "\n")
	board := Board{rows: createEmptyCounts(), cols: createEmptyCounts()}

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

func getBoards(inputs []string) []Board {
	boards := make([]Board, len(inputs))
	wg := sync.WaitGroup{}

	for i, input := range inputs {
		wg.Add(1)
		go func(i int, input string) {
			defer wg.Done()
			newBoard := newBoard(input)
			boards[i] = newBoard
		}(i, input)
	}
	wg.Wait()
	return boards
}
