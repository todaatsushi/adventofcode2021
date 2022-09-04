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

func ReadInput(filename string) [][2][]string {
	content, _ := os.ReadFile(filename)

	strContent := string(content)
	strContent = strings.TrimSpace(strContent)

	rawEntries := strings.Split(strContent, "\n")
	entries := make([][2][]string, len(rawEntries))

	var input []string
	var output []string
	for i, e := range rawEntries {
		e = strings.TrimSpace(e)
		input, output = readEntry(e)

		entries[i][0] = input
		entries[i][1] = output
	}
	return entries
}
