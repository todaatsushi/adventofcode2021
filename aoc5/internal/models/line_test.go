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
		if line.isReversed == true {
			t.Fatal("Line should not be marked as reversed")
		}
	})

	t.Run("Test new horizontal line", func(t *testing.T) {
		start := &Coordinate{x: 0, y: 10}
		end := &Coordinate{x: 10, y: 10}

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
		if line.isReversed == true {
			t.Fatal("Line should not be marked as reversed")
		}
	})

	t.Run("Test new diagonal line", func(t *testing.T) {
		start := &Coordinate{x: 1, y: 1}
		end := &Coordinate{x: 3, y: 3}

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
		if line.isReversed == true {
			t.Fatal("Line should not be marked as reversed")
		}
	})

	t.Run("Test new reverse diagonal line", func(t *testing.T) {
		start := &Coordinate{x: 3, y: 3}
		end := &Coordinate{x: 1, y: 1}

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
		if line.isReversed == false {
			t.Fatal("Line should be marked as reversed")
		}
	})

	t.Run("Test new line reverse horizontal line", func(t *testing.T) {
		start := &Coordinate{x: 0, y: 10}
		end := &Coordinate{x: 0, y: 0}

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
		if line.isReversed == false {
			t.Fatal("Line should be marked as reversed")
		}
	})

	t.Run("Test new line reverse vertical line", func(t *testing.T) {
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
		if line.isReversed == false {
			t.Fatal("Line should be marked as reversed")
		}
	})

	t.Run("Test is straight method", func(t *testing.T) {
		startStraight := &Coordinate{x: 10, y: 10}
		endStraight := &Coordinate{x: 0, y: 10}

		startNotStraight := &Coordinate{x: 10, y: 100}
		endNotStraight := &Coordinate{x: 1, y: 10}

		straightLine := newLine(startStraight, endStraight)
		notStraightLine := newLine(startNotStraight, endNotStraight)

		if straightLine.isStraight() == false {
			t.Fatal("Line should be evaluated as straight")
		}
		if notStraightLine.isStraight() == true {
			t.Fatal("Line shouldn't be evaluated as straight")
		}

	})
}
