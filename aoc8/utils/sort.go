package utils

import (
	"sort"
	"strings"
)

func SortString(s string) string {
	splitString := strings.Split(s, "")
	sort.Strings(splitString)
	sortedString := strings.Join(splitString, "")
	return sortedString
}
