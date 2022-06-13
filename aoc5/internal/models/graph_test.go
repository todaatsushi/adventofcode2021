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
		if x != 9 {
			t.Fatal("Expected 9 for max X, got", x)
		}
		if y != 9 {
			t.Fatal("Expected 9 for max Y, got", y)
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

		for r := 0; r < 9; r++ {
			if len((*graph)[r]) != 9 {
				t.Fatal("Graph row should have 9 cols in it, got:", len((*graph)[r]))
			}
		}

		if len(*graph) != 9 {
			t.Fatal("Graph should have 9 rows in it, got:", len(*graph))
		}

	})
}

func TestCreateGraph(t *testing.T) {
}
