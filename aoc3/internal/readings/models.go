package readings

import "log"

type Reading struct {
	value []int32
}

func (reading *Reading) pop() int32 {
	if len(reading.value) == 0 {
		log.Fatal("No values to pop")
	}
	first, remaining := reading.value[0], reading.value[1:]
	reading.value = remaining
	return first
}

func (reading *Reading) push(val int32) {
	reading.value = append(reading.value, val)
}
