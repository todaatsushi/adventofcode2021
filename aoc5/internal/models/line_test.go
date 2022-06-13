package models

import "testing"

func TestLine(t *testing.T) {
	t.Run("Test read lines from input", func(t *testing.T) {
		input := []string{
			"0,9 -> 5,9",
			"8,0 -> 0,8",
			"8,0 -> 0,8",
			"9,4 -> 3,4",
		}
		lines := ReadLines(input)

		if len(lines) != 4 {
			t.Fatal("Unexpected number of lines:", len(lines))
		}
	})

	t.Run("Test new vertical line", func(t *testing.T) {
		start := &Coordinate{x: 0, y: 0}
		end := &Coordinate{x: 0, y: 10}

		line := newLine(start, end)

		if line.start != start {
			t.Fatal("Line start coord is not the start coord")
		}
		if line.end != end {
			t.Fatal("Line start coord is not the start coord")
		}

		if line.isVertical == false {
			t.Fatal("Line should be marked as vertical")
		}
	})

	t.Run("Test new horizontal line", func(t *testing.T) {
		start := &Coordinate{x: 10, y: 10}
		end := &Coordinate{x: 0, y: 10}

		line := newLine(start, end)

		if line.start != start {
			t.Fatal("Line start coord is not the start coord")
		}
		if line.end != end {
			t.Fatal("Line start coord is not the start coord")
		}

		if line.isVertical == true {
			t.Fatal("Line should not be marked as vertical")
		}
	})
}
