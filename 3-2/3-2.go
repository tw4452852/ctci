package ctci

import (
	"fmt"
)

// How would you design a stack which, in addition to push and pop,
// also has a function min which returns the minimum element?
// Push, pop and min should all operate in O(1) time.

var (
	FullErr  = fmt.Errorf("stack full")
	EmptyErr = fmt.Errorf("stack empty")
)

type MinStack struct {
	data []byte
	mins []byte
	top  int
}

func NewMinStack(size int) *MinStack {
	return &MinStack{
		data: make([]byte, size),
		mins: make([]byte, size),
		top:  0,
	}
}

func (ms *MinStack) Push(v byte) error {
	top := ms.top
	if top == len(ms.data) {
		return FullErr
	}
	ms.data[top] = v
	if top > 0 && ms.mins[top-1] < v {
		ms.mins[top] = ms.mins[top-1]
	} else {
		ms.mins[top] = v
	}
	ms.top++
	return nil
}

func (ms *MinStack) Pop() (byte, error) {
	if ms.top == 0 {
		return 0, EmptyErr
	}
	ms.top--
	return ms.data[ms.top], nil
}

func (ms *MinStack) Min() (byte, error) {
	top := ms.top
	if top == 0 {
		return 0, EmptyErr
	}
	return ms.mins[top-1], nil
}
