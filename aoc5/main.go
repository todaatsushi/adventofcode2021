package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	startX int
	startY int
	endX   int
	endY   int
}

func getXAndY(lineStr string) (int, int) {
	coords := strings.Split(lineStr, ",")

	strX := strings.TrimSpace(coords[0])
	strY := strings.TrimSpace(coords[1])
	intX, _ := strconv.Atoi(strX)
	intY, _ := strconv.Atoi(strY)

	return intX, intY
}

func newLine(inputs string) Line {
	parts := strings.Split(inputs, "->")
	startX, startY := getXAndY(parts[0])
	endX, endY := getXAndY(parts[1])
	return Line{startX: startX, startY: startY, endX: endX, endY: endY}
}

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
	for _, l := range input {
		fmt.Println(newLine(l))
	}
}
