package main

import "StructureData/linkedlist"

type QueueLinkedList struct {
	items linkedlist.LinkedList
}

func (q *QueueLinkedList) Enqueue(item interface{}) {
	q.items.PutTail(&linkedlist.Node{Data: item})
}

func (q *QueueLinkedList) Dequeue() (interface{}, bool) {
	tail, ok := q.items.GetTail()
	if !ok {
		return nil, false
	}

	q.items.DeleteTail()

	return tail.Data, true
}

func (q QueueLinkedList) ToString() string {
	return q.items.ToString()
}
