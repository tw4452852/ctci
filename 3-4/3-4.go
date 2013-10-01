package ctci

import (
	"container/list"
	"fmt"
)

// In the classic problem of the Towers of Hanoi,
// you have 3 rods and N disks of different sizes which can slide onto any tower.
// The puzzle starts with disks sorted in ascending order of size from top to bottom
// (e.g., each disk sits on top of an even larger one). You have the following constraints:
//
// Only one disk can be moved at a time.
// A disk is slid off the top of one rod onto the next rod.
// A disk can only be placed on top of a larger disk.
//
// Write a program to move the disks from the first rod to the last using Stacks

func HanoiRecursion(n, from, bridge, to int) string {
	if n == 1 {
		return fmt.Sprintf("move %d from %d to %d\n", n, from, to)
	}
	s := HanoiRecursion(n-1, from, to, bridge)
	s += fmt.Sprintf("move %d from %d to %d\n", n, from, to)
	s += HanoiRecursion(n-1, bridge, from, to)
	return s
}

type operation struct {
	begin, end, from, bridge, to int
}

type Stack struct {
	*list.List
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func (s *Stack) Push(op *operation) {
	s.List.PushBack(op)
}

func (s *Stack) Pop() *operation {
	t := s.List.Back()
	s.List.Remove(t)
	return t.Value.(*operation)
}

func (s *Stack) IsEmpty() bool {
	return s.List.Len() == 0
}

func HanoiStack(n, from, bridge, to int) string {
	s := ""
	stk := NewStack()
	stk.Push(&operation{1, n, from, bridge, to})
	for !stk.IsEmpty() {
		op := stk.Pop()
		if op.begin != op.end {
			stk.Push(&operation{1, op.end - 1, op.bridge, op.from, op.to})
			stk.Push(&operation{op.end, op.end, op.from, op.bridge, op.to})
			stk.Push(&operation{1, op.end - 1, op.from, op.to, op.bridge})
		} else {
			s += fmt.Sprintf("move %d from %d to %d\n", op.end, op.from, op.to)
		}
	}
	return s
}
