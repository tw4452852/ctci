package ctci

import (
	"fmt"
)

// Describe how you could use a single array to implement three stacks.

var (
	WrongNumErr = fmt.Errorf("stackNum must between 0 and 2")
	FullErr     = fmt.Errorf("stack full")
	EmptyErr    = fmt.Errorf("stack empty")
)

type Stack3 struct {
	data       []byte
	tops       []int
	perMaxSize int
}

func NewStack3(size int) *Stack3 {
	return &Stack3{
		data:       make([]byte, size*3),
		tops:       []int{0, 0, 0},
		perMaxSize: size,
	}
}

func (s *Stack3) Push(stackNum int, v byte) error {
	if err := checkStackNumValid(stackNum); err != nil {
		return err
	}
	top := s.tops[stackNum]
	if top == s.perMaxSize {
		return FullErr
	}
	s.data[stackNum*s.perMaxSize+top] = v
	s.tops[stackNum]++
	return nil
}

func (s *Stack3) Pop(stackNum int) (byte, error) {
	if err := checkStackNumValid(stackNum); err != nil {
		return 0, err
	}
	if s.tops[stackNum] == 0 {
		return 0, EmptyErr
	}
	s.tops[stackNum]--
	return s.data[stackNum*s.perMaxSize+s.tops[stackNum]], nil
}

func checkStackNumValid(stackNum int) error {
	if stackNum < 0 || stackNum > 2 {
		return WrongNumErr
	}
	return nil
}

func (s *Stack3) String() string {
	str := "["
	for _, b := range s.data {
		str += fmt.Sprintf("%d ", b)
	}
	str += "]"
	return str
}
