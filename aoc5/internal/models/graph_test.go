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

func TestCreateGraph(t *testing.T) {
}
