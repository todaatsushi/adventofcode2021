package utils

import (
	"log"
	"os"
	"strings"
)

func readFile(filepath string) string {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Couldn't read file! Path: ", filepath)
	}

	contentStr := string(content)
	contentStr = strings.TrimSpace(contentStr)
	return contentStr
}
