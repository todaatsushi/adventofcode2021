package utils

func NumCharsMatches(s1, s2 string) int {
	charMap := make(map[rune]bool, len(s1))
	for _, rune := range s1 {
		charMap[rune] = true
	}

	matches := 0
	for _, rune := range s2 {
		if _, ok := charMap[rune]; ok {
			matches++
		}
	}
	return matches
}
