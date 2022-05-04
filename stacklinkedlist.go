package main

import "StructureData/linkedlist"

type StackLinkedList struct {
	items linkedlist.LinkedList
}

func (s *StackLinkedList) Push(item interface{}) {
	s.items.PutTail(&linkedlist.Node{Data: item})
}

func (s *StackLinkedList) Pop() (interface{}, bool) {
	head, ok := s.items.GetHead()
	if !ok {
		return nil, false
	}

	s.items.DeleteHead()

	return head.Data, true
}

func (s StackLinkedList) ToString() string {
	return s.items.ToString()
}
