package readings

import (
	"log"
	"os"
	"strings"
)

func ReadRawReadings(filename string) []Reading {
	content, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal("Couldn't read file:", err)
	}
	readings := string(content)
	readings = strings.TrimSpace(readings)
	splitReadings := strings.Split(readings, "\n")
	readingSlice := make([]Reading, 0)

	for _, reading := range splitReadings {
		newReading := Reading{}
		var val int64
		for _, char := range reading {
			if int32(char) == 48 {
				val = 0
			} else {
				val = 1
			}
			newReading.value = append(newReading.value, val)
		}
		readingSlice = append(readingSlice, newReading)
	}
	return readingSlice
}
