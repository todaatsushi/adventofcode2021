package utils

import (
	"os"
	"strings"
)

func parse(content string, size int) []string {
	content = strings.TrimSpace(content)
	splitRawOutput := strings.Split(content, " ")

	output := make([]string, size)
	for i, o := range splitRawOutput {
		output[i] = strings.TrimSpace(o)
	}
	return output
}

func readEntry(entry string) ([]string, []string) {
	splitContent := strings.Split(entry, "|")

	// Input section
	input := parse(splitContent[0], 10)

	// Output section
	output := parse(splitContent[1], 4)
	return input, output
}

func ReadInput(filename string) [][]string {
	content, _ := os.ReadFile(filename)

	strContent := string(content)
	strContent = strings.TrimSpace(strContent)

	rawEntries := strings.Split(strContent, "\n")
	entries := make([][]string, len(rawEntries))

	for i, e := range rawEntries {
		e = strings.TrimSpace(e)
		_, entries[i] = readEntry(e)
	}
	return entries
}
