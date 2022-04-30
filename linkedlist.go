package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l LinkedList) GetHead() (Node, bool) {
	if l.head == nil {
		return Node{}, false
	}

	return *l.head, true
}

func (l LinkedList) GetTail() (Node, bool) {
	if l.head == nil {
		return Node{}, false
	}

	node := l.head
	tail := l.head
	for l.length != 0 {
		tail = node
		node = node.next
		l.length--
	}

	return *tail, true
}

func (l *LinkedList) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *LinkedList) Postponed(n *Node) {
	head := l.head
	for head.next != nil {
		head = head.next
	}

	head.next = n
	l.length++
}

func (l *LinkedList) SearchWithValue(value interface{}) (Node, bool) {
	head := l.head
	for head != nil {
		if head.data == value {
			return *head, true
		}

		head = head.next
	}

	return Node{}, false
}

func (l *LinkedList) DeleteWithValue(value interface{}) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l *LinkedList) DeleteHead() {
	if l.length == 0 {
		return
	}

	l.head = l.head.next
	l.length--
}

func (l *LinkedList) DeleteTail() {
	if l.length == 0 {
		return
	}

	tail := l.head
	for tail.next != nil {
		tail = tail.next
	}

	tail = nil
	l.length--
}

func (l LinkedList) ToString() string {
	var result []string
	node := l.head
	for l.length != 0 {
		result = append(result, fmt.Sprintf("%v", node.data))
		node = node.next
		l.length--
	}

	return strings.Join(result, " ")
}

func printList() {
	myList := LinkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 16}
	node4 := &Node{data: 20}
	node5 := &Node{data: 22}
	node6 := &Node{data: 23}
	myList.Prepend(node1)
	myList.Prepend(node2)
	myList.Prepend(node3)
	myList.Postponed(node4)
	myList.Postponed(node5)
	myList.Prepend(node6)
	fmt.Println(myList.ToString())
	myList.DeleteHead()
	fmt.Println(myList.ToString())
	myList.DeleteTail()
	fmt.Println(myList.ToString())
	myList.DeleteWithValue(18)
	fmt.Println(myList.ToString())
	head, ok := myList.GetHead()
	if ok {
		fmt.Println(head.data)
	}
	tail, ok := myList.GetTail()
	if ok {
		fmt.Println(tail.data)
	}
}
