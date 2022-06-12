package models

import "sync"

type Line struct {
	start *Coordinate
	end   *Coordinate
}

func ReadLines(input []string) []*Line {
	lines := make([]*Line, len(input))
	wg := sync.WaitGroup{}

	for i, coords := range input {
		wg.Add(1)
		go func(lines []*Line, i int, coords string) {
			defer wg.Done()
			start, end := readCoordinates(coords)
			lines[i] = &Line{start: start, end: end}
		}(lines, i, coords)
	}

	wg.Wait()
	return lines
}
