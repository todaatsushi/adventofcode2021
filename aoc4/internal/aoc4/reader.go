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
	fmt.Println(splitContent, len(splitContent))
	for _, c := range splitContent {
		fmt.Println(c)
		fmt.Println()
	}
}
