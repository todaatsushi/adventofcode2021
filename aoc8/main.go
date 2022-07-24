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

func getCharCountMap() map[int]int {
	charCounts := map[int]int{
		2: 1,
		4: 4,
		3: 7,
		7: 8,
	}
	return charCounts
}

func main() {
	filename := os.Args[1]
	outputs := readInput(filename)
	charCounts := getCharCountMap()
	tally := map[int]int{1: 0, 4: 0, 7: 0, 8: 0}

	for _, outputSet := range outputs {
		for _, output := range outputSet {
			if num, ok := charCounts[len(output)]; ok {
				tally[num]++
			}
		}
	}
	total := 0
	for _, count := range tally {
		total += count
	}
	fmt.Println(total)
}
