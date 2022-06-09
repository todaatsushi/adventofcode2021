package aoc4

import (
	"log"
	"os"
	"strings"
)

func ReadInput(file string) (string, []string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal("Couldn't read file")
	}
	bingoContent := string(content)
	bingoContent = strings.TrimSpace(bingoContent)
	splitContent := strings.Split(bingoContent, "\n\n")
	return splitContent[0], splitContent[1:]
}

func PlayBingo(callsString string, boardsStrings []string) {
	bingoCalls := getCallQueue(callsString)
	boards := getBoards(boardsStrings)

	makeCalls(boards, bingoCalls)
}
