package linkedlist

import "fmt"

func (l *LinkedList) DeleteWithValue(value interface{}) {
	if l.size == 0 {
		return
	}

	if l.head.Data == value {
		l.DeleteHead()
		return
	}

	if l.tail.Data == value {
		l.DeleteTail()
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.Data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.size--
}

func (l *LinkedList) DeleteHead() {
	if l.size == 0 {
		return
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size--
		return
	}

	l.head = l.head.next
	l.size--
}

func (l *LinkedList) DeleteTail() {
	if l.size == 0 {
		return
	}

	if l.size == 1 {
		l.head = nil
		l.tail = nil
		l.size--
		return
	}

	newTail := l.head
	tailForDelete := l.head
	for tailForDelete.next != nil {
		newTail = tailForDelete
		tailForDelete = tailForDelete.next
	}

	tailForDelete = nil
	l.tail = newTail
	l.size--
}

func (l *LinkedList) DeleteIndex(index int) error {
	if l.size <= index {
		return fmt.Errorf("index is greater than length of list")
	}

	if index < 0 {
		return fmt.Errorf("the index must be greater than 0")
	}

	if index == 0 {
		l.DeleteHead()
		return nil
	}

	if index == l.size-1 {
		l.DeleteTail()
		return nil
	}

	previous := l.head
	for i := 0; i < index-1; i++ {
		previous = previous.next
	}

	previous.next = previous.next.next
	l.size--

	return nil
}
