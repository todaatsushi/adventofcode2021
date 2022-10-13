package grid

import (
	"log"
	"strconv"
	"strings"
)

func getCoordinateModifiers() [4][2]int {
	var moves [4][2]int

	moves[0] = [2]int{1, 0}
	moves[1] = [2]int{-1, 0}
	moves[2] = [2]int{0, 1}
	moves[3] = [2]int{0, -1}

	return moves
}

type Grid struct {
	points [][]int
}

func NewGrid(fileLines []string) Grid {
	numRows := len(fileLines)
	numCols := len(fileLines[0])

	points := make([][]int, numRows)

	for rowNum, strRow := range fileLines {
		colSlice := strings.Split(strRow, "")
		row := make([]int, numCols)

		for colNum, col := range colSlice {
			val, e := strconv.Atoi(col)
			if e != nil {
				log.Fatal("Couldn't convert val to int:", col)
			}
			row[colNum] = val
		}

		points[rowNum] = row
	}

	g := Grid{points: points}

	return g
}
