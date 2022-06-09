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
