package linkedlist

import "fmt"

func (l *LinkedList) SearchWithValue(value interface{}) (Node, bool) {
	head := l.head
	for head != nil {
		if head.Data == value {
			return *head, true
		}

		head = head.next
	}

	return Node{}, false
}

func (l *LinkedList) GetIndex(index int) (Node, error) {
	if l.length <= index {
		return Node{}, fmt.Errorf("index is greater than length of list")
	}

	if index < 0 {
		return Node{}, fmt.Errorf("the index must be greater than 0")
	}

	if index == 0 {
		n, ok := l.GetHead()
		if !ok {
			return Node{}, fmt.Errorf("index not found")
		}

		return n, nil
	}

	if index == l.length-1 {
		n, ok := l.GetTail()
		if !ok {
			return Node{}, fmt.Errorf("index not found")
		}

		return n, nil
	}

	previous := l.head
	for i := 0; i < index; i++ {
		previous = previous.next
	}

	return *previous, nil
}

func (l LinkedList) GetHead() (Node, bool) {
	if l.head == nil {
		return Node{}, false
	}

	return *l.head, true
}

func (l LinkedList) GetTail() (Node, bool) {
	if l.tail == nil {
		return Node{}, false
	}

	return *l.tail, true
}
