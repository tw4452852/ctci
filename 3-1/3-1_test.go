package ctci

import (
	"testing"
)

func TestStack3(t *testing.T) {
	stk := NewStack3(2)
	err := stk.Push(3, 1)
	if err != WrongNumErr {
		t.Errorf("expect err %v, but got err %v\n", WrongNumErr, err)
	}
	err = stk.Push(0, 1)
	if err != nil {
		t.Errorf("expect err %v, but got err %v\n", nil, err)
	}
	err = stk.Push(0, 2)
	if err != nil {
		t.Errorf("expect err %v, but got err %v\n", nil, err)
	}
	err = stk.Push(0, 3)
	if err != FullErr {
		t.Errorf("expect err %v, but got err %v\n", FullErr, err)
	}
	err = stk.Push(1, 3)
	if err != nil {
		t.Errorf("expect err %v, but got err %v\n", nil, err)
	}
	if s := stk.String(); s != "[1 2 3 0 0 0 ]" {
		t.Errorf("data %q isn't the expect one\n", s)
	}
	var v byte
	v, err = stk.Pop(0)
	if err != nil {
		t.Errorf("expect err %v, but got err %v\n", nil, err)
	}
	if v != 2 {
		t.Errorf("expect %v, but got %v\n", 2, v)
	}
	v, err = stk.Pop(0)
	if err != nil {
		t.Errorf("expect err %v, but got err %v\n", nil, err)
	}
	if v != 1 {
		t.Errorf("expect %v, but got %v\n", 1, v)
	}
	v, err = stk.Pop(0)
	if err != EmptyErr {
		t.Errorf("expect err %v, but got err %v\n", EmptyErr, err)
	}
}
