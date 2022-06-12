package aoc5

import (
	"log"
	"os"
	"strings"
)

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Couldn't read file:", filename)
	}
	rawInput := string(content)
	rawInput = strings.TrimSpace(rawInput)
	input := strings.Split(rawInput, "\n")
	return input
}
