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

func GetTally(lines []Line) int {
	tally := 0
	graph := *NewGraph(lines)

	for _, l := range lines {
		if l.IsStraight() == true {
			tallyStraightLine(&l, &tally, &graph)
		}
	}
	return tally
}
