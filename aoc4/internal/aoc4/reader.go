package aoc4

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInput(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Couldn't read file")
	}
	bingoContent := string(content)
	bingoContent = strings.TrimSpace(bingoContent)
	splitContent := strings.Split(bingoContent, "\n\n")

	bingoCalls := getCallQueue(splitContent[0])
	rawBoardInputs := splitContent[1:]
	boards := getBoards(rawBoardInputs)

	fmt.Println("Got calls:", bingoCalls)
	fmt.Println("Boards:", boards)
}
