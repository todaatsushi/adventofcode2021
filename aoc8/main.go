package main

import (
	"fmt"
	"os"
	"strings"
)

func readEntry(entry string) []string {
	splitContent := strings.Split(entry, "|")

	rawOutput := strings.TrimSpace(splitContent[1])
	splitRawOutput := strings.Split(rawOutput, " ")

	output := make([]string, 4)
	for i, o := range splitRawOutput {
		output[i] = strings.TrimSpace(o)
	}
	return output

}

func readInput(filename string) [][]string {
	content, _ := os.ReadFile(filename)

	strContent := string(content)
	strContent = strings.TrimSpace(strContent)

	rawEntries := strings.Split(strContent, "\n")
	entries := make([][]string, len(rawEntries))

	for i, e := range rawEntries {
		e = strings.TrimSpace(e)
		entries[i] = readEntry(e)
	}
	return entries
}

func main() {
	filename := os.Args[1]
	output := readInput(filename)
	fmt.Println(output)
}
