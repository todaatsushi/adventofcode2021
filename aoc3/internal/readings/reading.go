package readings

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func fmtReadingValue(readingValue []int64) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(readingValue)), ""), "[]")
}

func convertReadingString(readingValueStr string) int64 {
	intValue, _ := strconv.ParseInt(readingValueStr, 2, 64)
	return intValue
}

type Reading struct {
	value []int64
}

func (reading *Reading) pop() int64 {
	if len(reading.value) == 0 {
		log.Fatal("No values to pop")
	}
	first, remaining := reading.value[0], reading.value[1:]
	reading.value = remaining
	return first
}

func (reading *Reading) push(val int64) {
	reading.value = append(reading.value, val)
}

func (reading *Reading) fmtValue() string {
	return fmtReadingValue(reading.value)
}

func (reading *Reading) convertToInt() int64 {
	stringValue := reading.fmtValue()
	return convertReadingString(stringValue)
}
