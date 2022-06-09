package aoc4

import "sync"

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

func checkBoards(toCheck int, boards []*Board) {
	wg := sync.WaitGroup{}
	for _, board := range boards {
		wg.Add(1)
		go func(board *Board, val int) {
			defer wg.Done()
			board.check(val)
		}(board, toCheck)
	}
	wg.Wait()
}

func makeCalls(boards []*Board, calls chan int) {
	for call := range calls {
		checkBoards(call, boards)
	}
}
