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
		readingSlice = append(readingSlice, Reading{value: reading})
	}
	return readingSlice
}
