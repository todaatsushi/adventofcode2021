package main

import (
	"adventofcode2021/aoc5/helpers"
	"fmt"
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

func main() {
	filename := os.Args[1]

	input := readInput(filename)
	lines := helpers.ReadLines(input)
	tally := helpers.GetTally(lines)

	fmt.Println("Tally:", tally)
}
