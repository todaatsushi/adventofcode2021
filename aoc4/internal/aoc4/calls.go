package aoc4

import (
	"strconv"
	"strings"
)

func getCallQueue(inputs string) chan int {
	calls := strings.Split(inputs, ",")
	callChannel := make(chan int, len(calls))

	var callInt int
	var err error
	for _, call := range calls {
		call = strings.TrimSpace(call)
		callInt, err = strconv.Atoi(call)
		if err == nil {
			callChannel <- callInt
		}
	}
	return callChannel
}
