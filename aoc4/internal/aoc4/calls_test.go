package aoc4

import (
	"testing"
)

func TestGetCallsQueue(t *testing.T) {
	input := "7,4,9,5,11,"
	expected := make(chan int, 5)
	toAdd := []int{7, 4, 9, 5, 11}
	for _, i := range toAdd {
		expected <- i
	}
	actual := getCallQueue(input)

	var actualVar int
	var expectedVar int
	for x := 0; x < 5; x++ {
		actualVar = <-actual
		expectedVar = <-expected

		if actualVar != expectedVar {
			t.Fatal("Values aren't the same:", actualVar, expectedVar)
		}
	}
}
