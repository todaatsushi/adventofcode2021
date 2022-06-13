package models

import (
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
		graph := CreateGraph(lines)

		for r := 0; r < 10; r++ {
			if len((*graph)[r]) != 10 {
				t.Fatal("Graph row should have 10 cols in it, got:", len((*graph)[r]))
			}
		}

		if len(*graph) != 10 {
			t.Fatal("Graph should have 10 rows in it, got:", len(*graph))
		}

		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				if (*graph)[r][c] != 0 {
					t.Fatal("Slice value should init at 0, got:", (*graph)[r][c])
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
		graph := CreateGraph(lines)
		ReadLinesToGraph(graph, lines)

		for i := 0; i < 9; i++ {
			if (*graph)[0][i] != 1 {
				t.Fatalf("Expected init value of %d at %d,%d. Got %v", 1, 0, i, (*graph)[0][i])
			}
		}
	})

	t.Run("Test draw horizontal line", func(t *testing.T) {
		input := []string{
			"0,0 -> 9,0",
		}
		lines := ReadLines(input)
		graph := CreateGraph(lines)
		ReadLinesToGraph(graph, lines)

		for i := 0; i < 9; i++ {
			if (*graph)[i][0] != 1 {
				t.Fatalf("Expected init value of %d at %d,%d. Got %v", 1, 0, i, (*graph)[0][i])
			}
		}

	})

	t.Run("Test draw multiple lines", func(t *testing.T) {
		input := []string{
			"0,0 -> 0,9",
			"0,0 -> 9,0",
		}
		lines := ReadLines(input)
		graph := CreateGraph(lines)
		ReadLinesToGraph(graph, lines)

		expected := [10][10]int{
			{2, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}
		for r := 0; r < 10; r++ {
			for c := 0; c < 10; c++ {
				if (*graph)[r][c] != expected[r][c] {
					t.Fatalf("Graph val @(%d, %d) different - got %d, expected %d", r, c, (*graph)[r][c], expected[r][c])
				}
			}
		}
	})
}
