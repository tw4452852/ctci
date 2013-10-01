package ctci

import (
	"container/list"
	"fmt"
)

// Implement a MyQueue class which implements a queue using two stacks.

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

func (s *Stack) IsEmpty() bool {
	return s.List.Len() == 0
}

var EmptyErr = fmt.Errorf("empty queue")

type MyQueue struct {
	en, de *Stack
}

func NewMyQueue() *MyQueue {
	return &MyQueue{NewStack(), NewStack()}
}

func (q *MyQueue) EnQueue(v interface{}) {
	for !q.de.IsEmpty() {
		q.en.Push(q.de.Pop())
	}
	q.en.Push(v)
}

func (q *MyQueue) DeQueue() (interface{}, error) {
	if q.de.IsEmpty() && q.en.IsEmpty() {
		return nil, EmptyErr
	}
	for !q.en.IsEmpty() {
		q.de.Push(q.en.Pop())
	}
	return q.de.Pop(), nil
}
