package utils

import (
	"log"
	"os"
	"strconv"
)

func GetArgs() (string, int, bool) {
	filepath := os.Args[1]

	numSteps, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Couldn't convert numSteps to an int. Got:", os.Args[2])
	}

	display := os.Args[3] == "yes"

	return filepath, numSteps, display
}
