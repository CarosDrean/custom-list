package main

type QueueLinkedList struct {
	items LinkedList
}

func (q *QueueLinkedList) Enqueue(item interface{}) {
	q.items.Postponed(&Node{data: item})
}

func (q *QueueLinkedList) Dequeue() (interface{}, bool) {
	tail, ok := q.items.GetTail()
	if !ok {
		return nil, false
	}

	q.items.DeleteTail()

	return tail.data, true
}
