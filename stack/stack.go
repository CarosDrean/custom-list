package stack

import "fmt"

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.items) == 0 {
		return nil, false
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item, true
}

func (s Stack) ToString() string {
	return fmt.Sprintf("%v", s.items)
}
