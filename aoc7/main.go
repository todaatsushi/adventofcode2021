package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func parseCrabsFromFile(filename string) []int {
	content, _ := os.ReadFile(filename)
	rawContent := string(content)
	positions := strings.Split(rawContent, ",")

	crabs := make([]int, len(positions))
	wg := sync.WaitGroup{}
	for i, p := range positions {
		wg.Add(1)
		p = strings.TrimSpace(p)
		go func(crabs *[]int, index int, val string) {
			defer wg.Done()
			asNum, _ := strconv.Atoi(val)
			(*crabs)[index] = asNum
		}(&crabs, i, p)
	}
	wg.Wait()

	sort.Ints(crabs)
	return crabs
}

func main() {
	filename := os.Args[1]
	crabs := parseCrabsFromFile(filename)

	fmt.Println(crabs)
}
