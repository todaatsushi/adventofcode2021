package readings

func getMostCommonValueAtPosition(readings []Reading, position int) int64 {
	tally := Tally{one: 0, zero: 0}
	tally.tallyReadingsAtPosition(readings, position)

	max, err := tally.getMax()
	if err != nil {
		return 1
	}
	return max
}
