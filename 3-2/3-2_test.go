package ctci

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	stk := NewMinStack(2)
	err := stk.Push(1)
	if err != nil {
		t.Errorf("expect err %v, but got err %v\n", nil, err)
	}
	err = stk.Push(2)
	if err != nil {
		t.Errorf("expect err %v, but got err %v\n", nil, err)
	}
	err = stk.Push(3)
	if err != FullErr {
		t.Errorf("expect err %v, but got err %v\n", FullErr, err)
	}
	var v byte
	v, err = stk.Min()
	if err != nil {
		t.Errorf("expect err %v, but got err %v\n", nil, err)
	}
	if v != 1 {
		t.Errorf("expect min %d, but got min %d\n", 1, v)
	}
	v, err = stk.Pop()
	if err != nil {
		t.Errorf("expect err %v, but got err %v\n", nil, err)
	}
	if v != 2 {
		t.Errorf("expect pop value %d, but got %d\n", 2, v)
	}
	v, err = stk.Pop()
	if err != nil {
		t.Errorf("expect err %v, but got err %v\n", nil, err)
	}
	if v != 1 {
		t.Errorf("expect pop value %d, but got %d\n", 1, v)
	}
	v, err = stk.Pop()
	if err != EmptyErr {
		t.Errorf("expect err %v, but got err %v\n", EmptyErr, err)
	}
	if v != 0 {
		t.Errorf("expect pop value %d, but got %d\n", 0, v)
	}
}
