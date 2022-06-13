package models

import (
	"fmt"
	"testing"
)

func TestGraphHelpers(t *testing.T) {
	t.Run("Test get graph size", func(t *testing.T) {
		input := []string{
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
		}
		lines := ReadLines(input)

		x, y := getGraphSize(lines)
		if x != 10 {
			t.Fatal("Expected 10 for max X, got", x)
		}
		if y != 10 {
			t.Fatal("Expected 10 for max Y, got", y)
		}
	})

	t.Run("Test create new graph", func(t *testing.T) {
		input := []string{
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
		}
		lines := ReadLines(input)
		graph := NewGraph(lines)

		for r := 0; r < 10; r++ {
			if len((*graph).points[r]) != 10 {
				t.Fatal("Graph row should have 10 cols in it, got:", len((*graph).points[r]))
			}
		}

		if len((*graph).points) != 10 {
			t.Fatal("Graph should have 10 rows in it, got:", len((*graph).points))
		}

		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				if (*graph).points[r][c] != 0 {
					t.Fatal("Slice value should init at 0, got:", (*graph).points[r][c])
				}
			}
		}

	})
}

func TestGraphReading(t *testing.T) {
	t.Run("Test draw vertical line", func(t *testing.T) {
		input := []string{
			"0,0 -> 0,9",
		}
		lines := ReadLines(input)
		graph := NewGraph(lines)
		graph.ReadLines()

		for i := 0; i < 9; i++ {
			if (*graph).points[0][i] != 1 {
				t.Fatalf("Expected init value of %d at %d,%d. Got %v", 1, 0, i, (*graph).points[0][i])
			}
		}
	})

	t.Run("Test draw horizontal line", func(t *testing.T) {
		input := []string{
			"0,0 -> 9,0",
		}
		lines := ReadLines(input)
		graph := NewGraph(lines)
		graph.ReadLines()

		for i := 0; i < 9; i++ {
			if (*graph).points[i][0] != 1 {
				t.Fatalf("Expected init value of %d at %d,%d. Got %v", 1, 0, i, (*graph).points[0][i])
			}
		}

	})

	t.Run("Test draw backwards vertical line", func(t *testing.T) {
		input := []string{
			"0,9 -> 0,0",
		}
		lines := ReadLines(input)
		graph := NewGraph(lines)
		graph.ReadLines()

		for i := 0; i < 9; i++ {
			if (*graph).points[0][i] != 1 {
				t.Fatalf("Expected init value of %d at %d,%d. Got %v", 1, 0, i, (*graph).points[0][i])
			}
		}
	})

	t.Run("Test draw backwards horizontal line", func(t *testing.T) {
		input := []string{
			"9,0 -> 0,0",
		}
		lines := ReadLines(input)
		graph := NewGraph(lines)
		graph.ReadLines()

		for i := 0; i < 9; i++ {
			if (*graph).points[i][0] != 1 {
				t.Fatalf("Expected init value of %d at %d,%d. Got %v", 1, 0, i, (*graph).points[0][i])
			}
		}
	})

	t.Run("Test draw multiple lines", func(t *testing.T) {
		input := []string{
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
			"2,2 -> 2,1",
			"7,0 -> 7,4",
			"6,4 -> 2,0",
			"0,9 -> 2,9",
			"3,4 -> 1,4",
			"0,0 -> 8,8",
			"5,5 -> 8,2",
		}
		lines := ReadLines(input)
		graph := NewGraph(lines)
		graph.ReadLines()

		expected := [10][10]int{
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{2, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		}
		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				if (*graph).points[r][c] != expected[r][c] {
					fmt.Println("Graph incorrect:")
					graph.Describe(true)
					t.Fatalf("Graph val @(%d, %d) different - got %d, expected %d", r, c, (*graph).points[r][c], expected[r][c])
				}
			}
		}
	})

	t.Run("Test get tallies from graph after read", func(t *testing.T) {
		input := []string{
			"0,0 -> 0,9",
			"9,0 -> 0,0",
		}
		lines := ReadLines(input)
		graph := NewGraph(lines)
		graph.ReadLines()

		pointsOver2 := 1
		pointsOver1 := 19
		pointsOver10 := 0
		vals := []int{2, 1, 10}

		expectedCounts := []int{pointsOver2, pointsOver1, pointsOver10}
		actualCounts := make([]int, 3)

		var c int
		for i, v := range vals {
			c = graph.PointsWithOverXNumberOfLines(v)
			actualCounts[i] = c
		}

		for i, ans := range actualCounts {
			if ans != expectedCounts[i] {
				t.Fatalf("Tally for val %d incorrect - expected %d, got %d", vals[i], expectedCounts[i], ans)
			}
		}

	})
}
