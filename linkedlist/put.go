package linkedlist

import "fmt"

func (l *LinkedList) PutHead(n *Node) {
	if l.length == 0 {
		l.head = n
		l.tail = n
		l.length++
		return
	}

	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *LinkedList) PutTail(n *Node) {
	if l.length == 0 {
		l.PutHead(n)
		return
	}

	l.tail.next = n
	l.tail = n
	l.length++
}

func (l *LinkedList) PutIndex(index int, n *Node) error {
	if l.length < index {
		return fmt.Errorf("index is greater than length of list")
	}

	if index < 0 {
		return fmt.Errorf("the index must be greater than 0")
	}

	if index == 0 {
		l.PutHead(n)
		return nil
	}

	if index == l.length {
		l.PutTail(n)
		return nil
	}

	previous := l.head
	for i := 0; i < index-1; i++ {
		previous = previous.next
	}

	n.next = previous.next
	previous.next = n
	l.length++

	return nil
}
