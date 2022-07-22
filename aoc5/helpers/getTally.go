package helpers

func tallyStraightLine(line *Line, tally *int, graph *Graph) {
	var start, end int
	if line.StartX == line.EndX {
		start, end = line.GetStartAndEnd("y")
		for i := start; i <= end; i++ {
			graph.Points[line.StartX][i]++
			if graph.Points[line.StartX][i] == 2 {
				*tally++
			}
		}
	} else {
		start, end = line.GetStartAndEnd("x")
		for i := start; i <= end; i++ {
			graph.Points[i][line.StartY]++
			if graph.Points[i][line.StartY] == 2 {
				*tally++
			}
		}
	}
}

func tallyDiagonalLine(line *Line, tally *int, graph *Graph) {
	startX, endX := line.GetStartAndEnd("x")
	startY, endY := line.GetStartAndEnd("y")
	var movement int

	if endY < startY {
		movement = -1
	} else {
		movement = 1
	}

	x, y := startX, startY

	for {
		if x == endX+1 && y == endY+movement {
			break
		}
		graph.Points[x][y]++
		if graph.Points[x][y] == 2 {
			*tally++
		}

		x++
		y += movement
	}
}

func GetTally(lines []Line) int {
	tally := 0
	graph := *NewGraph(lines)

	for _, l := range lines {
		if l.IsStraight() == true {
			tallyStraightLine(&l, &tally, &graph)
		} else {
			tallyDiagonalLine(&l, &tally, &graph)
		}
	}
	return tally
}
