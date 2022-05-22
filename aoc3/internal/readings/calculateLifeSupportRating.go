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

func getLeastCommonValueAtPosition(readings []Reading, position int) int64 {
	tally := Tally{one: 0, zero: 0}
	tally.tallyReadingsAtPosition(readings, position)

	max, err := tally.getMin()
	if err != nil {
		return 0
	}
	return max
}

func GetO2Readings(readings []Reading, position int) Reading {
	if len(readings) == 1 {
		return readings[0]
	}
	o2Readings := make([]Reading, 0)
	filter := getMostCommonValueAtPosition(readings, position)

	for _, reading := range readings {
		if reading.value[position] == filter {
			o2Readings = append(o2Readings, reading)
		}
	}
	return GetO2Readings(o2Readings, position+1)
}
