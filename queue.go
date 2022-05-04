package main

import "fmt"

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, bool) {
	if len(q.items) == 0 {
		return nil, false
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, true
}

func (q Queue) ToString() string {
	return fmt.Sprintf("%v", q.items)
}
