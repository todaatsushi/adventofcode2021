package utils

type dict struct {
	Table map[string]int
	Index [10]string
}

func NewDict() *dict {
	var index [10]string
	table := make(map[string]int, 10)
	return &dict{Table: table, Index: index}
}

func (d *dict) MapVal(pattern string, val int) {
	d.Table[pattern] = val
	d.Index[val] = pattern
}
