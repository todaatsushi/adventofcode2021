package input

import (
	"log"
	"os"
	"strings"
)

func ReadInput(filepath string) []string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Couldn't read file:", filepath)
	}

	rawInput := string(content)
	rawInput = strings.TrimSpace(rawInput)

	input := strings.Split(rawInput, "\n")
	return input
}
