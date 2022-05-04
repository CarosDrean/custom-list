package linkedlist

import "fmt"

func (l *LinkedList) PutHead(n *Node) {
	if l.size == 0 {
		l.head = n
		l.tail = n
		l.size++
		return
	}

	second := l.head
	l.head = n
	l.head.next = second
	l.size++
}

func (l *LinkedList) PutTail(n *Node) {
	if l.size == 0 {
		l.PutHead(n)
		return
	}

	l.tail.next = n
	l.tail = n
	l.size++
}

func (l *LinkedList) PutIndex(index int, n *Node) error {
	if l.size < index {
		return fmt.Errorf("index is greater than length of list")
	}

	if index < 0 {
		return fmt.Errorf("the index must be greater than 0")
	}

	if index == 0 {
		l.PutHead(n)
		return nil
	}

	if index == l.size {
		l.PutTail(n)
		return nil
	}

	previous := l.head
	for i := 0; i < index-1; i++ {
		previous = previous.next
	}

	n.next = previous.next
	previous.next = n
	l.size++

	return nil
}
