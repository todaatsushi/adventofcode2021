package utils

import "errors"

type Queue struct {
	elems [][2]int
}

func (q *Queue) Add(elem [2]int) {
	q.elems = append(q.elems, elem)
}

func (q *Queue) Next() ([2]int, error) {
	if q.IsEmpty() == true {
		// Find better way of returning null val
		nullValue := [2]int{-1, -1}
		return nullValue, errors.New("Queue is empty!")
	}
	elem := q.elems[0]
	q.elems = q.elems[1:]
	return elem, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.elems) == 0
}

func NewQueue() Queue {
	elems := make([][2]int, 0)
	return Queue{elems: elems}
}
