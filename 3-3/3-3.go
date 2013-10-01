package ctci

import (
	"container/list"
	"fmt"
)

// Imagine a (literal) stack of plates.
// If the stack gets too high, it might topple.
// Therefore, in real life, we would likely start a new stack
// when the previous stack exceeds some threshold.
// Implement a data structure SetOfStacks that mimics this.
// SetOfStacks should be composed of several stacks,
// and should create a new stack once the previous one exceeds capacity.
// SetOfStacks.push() and SetOfStacks.pop() should behave identically to a single stack
// (that is, pop() should return the same values as it would if there were just a single stack).
//
// FOLLOW UP
//
// Implement a function popAt(int index) which performs a pop operation on a specific sub-stack.

var (
	EmptyErr = fmt.Errorf("stack empty")
	FullErr  = fmt.Errorf("stack full")
	IndexErr = fmt.Errorf("invalid index")
)

type Stack struct {
	l        *list.List
	capacity int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		l:        list.New(),
		capacity: capacity,
	}
}

func (s *Stack) Push(v byte) {
	s.l.PushBack(v)
}

func (s *Stack) Pop() byte {
	t := s.l.Back()
	s.l.Remove(t)
	return t.Value.(byte)
}

func (s *Stack) IsEmpty() bool {
	return s.l.Len() == 0
}

func (s *Stack) IsFull() bool {
	return s.l.Len() == s.capacity
}

type SetOfStacks struct {
	stacks   []*Stack
	capacity int
	cur      int
}

func NewSetOfStacks(capacity int) *SetOfStacks {
	return &SetOfStacks{
		stacks:   []*Stack{NewStack(capacity)},
		capacity: capacity,
		cur:      0,
	}
}

func (ss *SetOfStacks) Push(v byte) {
	if ss.stacks[ss.cur].IsFull() {
		if ss.cur == len(ss.stacks)-1 {
			ss.stacks = append(ss.stacks, NewStack(ss.capacity))
		}
		ss.cur++
	}
	ss.stacks[ss.cur].Push(v)
}

func (ss *SetOfStacks) Pop() (byte, error) {
	for ss.stacks[ss.cur].IsEmpty() && ss.cur > 0 {
		ss.cur--
	}
	if ss.cur == 0 && ss.stacks[ss.cur].IsEmpty() {
		return 0, EmptyErr
	}
	return ss.stacks[ss.cur].Pop(), nil

}

func (ss *SetOfStacks) PopAt(index int) (byte, error) {
	if index < 0 || index >= len(ss.stacks) {
		return 0, IndexErr
	}
	for ss.stacks[index].IsEmpty() && index > 0 {
		index--
	}
	if index == 0 && ss.stacks[index].IsEmpty() {
		return 0, EmptyErr
	}
	return ss.stacks[index].Pop(), nil
}
