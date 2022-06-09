package aoc4

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Board struct {
	nums     [5][5]int
	score    int
	rows     [5]int
	cols     [5]int
	complete bool
}

func (board *Board) check(number int) {
	var current int
	if board.complete == true {
		return
	}

	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			current = board.nums[r][c]
			if number == current {
				board.score -= current
				board.rows[r] += 1
				board.cols[c] += 1
				if board.rows[r] == 5 || board.cols[c] == 5 {
					board.complete = true
					board.score *= number
					fmt.Println("Board complete: score -", board.score)
					return
				}
			}
		}
	}
}

func createEmptyCounts() [5]int {
	return [5]int{0, 0, 0, 0, 0}
}

func newBoard(input string) *Board {
	rows := strings.Split(input, "\n")
	board := &Board{rows: createEmptyCounts(), cols: createEmptyCounts(), score: 0, complete: false}

	var vals []string
	var boardNum int
	var colNum int
	for rowNum, row := range rows {
		row = strings.TrimSpace(row)
		vals = strings.Split(row, " ")
		colNum = 0

		for _, val := range vals {
			numVal, err := strconv.Atoi(val)
			if err == nil {
				boardNum = numVal
				board.nums[rowNum][colNum] = boardNum
				board.score += boardNum
				colNum++
			}
		}
	}
	return board
}

func getBoards(inputs []string) []*Board {
	boards := make([]*Board, len(inputs))
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
