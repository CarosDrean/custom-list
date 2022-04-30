package main

type StackLinkedList struct {
	items LinkedList
}

func (s *StackLinkedList) Push(item interface{}) {
	s.items.Postponed(&Node{data: item})
}

func (s *StackLinkedList) Pop() (interface{}, bool) {
	head, ok := s.items.GetHead()
	if !ok {
		return nil, false
	}

	s.items.DeleteHead()

	return head.data, true
}
