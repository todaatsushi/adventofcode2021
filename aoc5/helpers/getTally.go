package helpers

func GetTally(lines []Line) int {
	tally := 0
	graph := *NewGraph(lines)

	var start int
	var end int
	for _, l := range lines {
		if l.IsStraight() == true {
			if l.StartX == l.EndX {
				start, end = l.GetStartAndEnd("y")
				for i := start; i <= end; i++ {
					graph.Points[l.StartX][i]++
					if graph.Points[l.StartX][i] == 2 {
						tally++
					}
				}
			} else {
				start, end = l.GetStartAndEnd("x")
				for i := start; i <= end; i++ {
					graph.Points[i][l.StartY]++
					if graph.Points[i][l.StartY] == 2 {
						tally++
					}
				}
			}
		}
	}
	return tally
}
