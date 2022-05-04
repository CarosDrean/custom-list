package queue

import "github.com/CarosDrean/custom-list/linkedlist"

type QueueLinkedList struct {
	items linkedlist.LinkedList
}

func (q *QueueLinkedList) Enqueue(item interface{}) {
	q.items.PutTail(&linkedlist.Node{Data: item})
}

func (q *QueueLinkedList) Dequeue() (interface{}, bool) {
	tail, ok := q.items.GetHead()
	if !ok {
		return nil, false
	}

	q.items.DeleteHead()

	return tail.Data, true
}

func (q QueueLinkedList) ToString() string {
	return q.items.ToString()
}
