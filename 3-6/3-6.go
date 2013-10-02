package ctci

import (
	"container/list"
)

// Write a program to sort a stack in ascending order.
// You should not make any assumptions about how the stack is implemented.
// The following are the only functions that should be used to write this program: push | pop | peek | isEmpty.

type Stack struct {
	*list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(e interface{}) {
	s.List.PushBack(e)
}

func (s *Stack) Pop() interface{} {
	t := s.List.Back()
	s.List.Remove(t)
	return t.Value
}

func (s *Stack) Peek() interface{} {
	return s.List.Back().Value
}

func (s *Stack) IsEmpty() bool {
	return s.List.Len() == 0
}

func SortStack(s *Stack) *Stack {
	tmpStk := NewStack()
	for !s.IsEmpty() {
		top := s.Pop()
		for !tmpStk.IsEmpty() && top.(int) < tmpStk.Peek().(int) {
			s.Push(tmpStk.Pop())
		}
		tmpStk.Push(top)
	}
	return tmpStk
}
