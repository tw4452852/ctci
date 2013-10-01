package ctci

import (
	"testing"
)

func TestSetOfStacks(t *testing.T) {
	ss := NewSetOfStacks(2)
	for i := 0; i < 5; i++ {
		ss.Push(byte(i))
	}
	if cur := ss.cur; cur != 2 {
		t.Errorf("expect cur %d, but got %d\n", 2, cur)
	}
	if cnt := len(ss.stacks); cnt != 3 {
		t.Errorf("expect stack num %d, but got %d\n", 3, cnt)
	}
	if v, e := ss.PopAt(4); e != IndexErr {
		t.Errorf("expect err %v, but got %v\n", IndexErr, e)
	} else if v != 0 {
		t.Errorf("expect v %v, but got %v\n", 0, v)
	}
	if v, e := ss.PopAt(1); e != nil {
		t.Errorf("expect err %v, but got %v\n", nil, e)
	} else if v != 3 {
		t.Errorf("expect v %v, but got %v\n", 3, v)
	}
	if v, e := ss.PopAt(1); e != nil {
		t.Errorf("expect err %v, but got %v\n", nil, e)
	} else if v != 2 {
		t.Errorf("expect v %v, but got %v\n", 2, v)
	}
	if v, e := ss.Pop(); e != nil {
		t.Errorf("expect err %v, but got %v\n", nil, e)
	} else if v != 4 {
		t.Errorf("expect v %v, but got %v\n", 4, v)
	}
	if v, e := ss.Pop(); e != nil {
		t.Errorf("expect err %v, but got %v\n", nil, e)
	} else if v != 1 {
		t.Errorf("expect v %v, but got %v\n", 1, v)
	}
	if v, e := ss.Pop(); e != nil {
		t.Errorf("expect err %v, but got %v\n", nil, e)
	} else if v != 0 {
		t.Errorf("expect v %v, but got %v\n", 0, v)
	}
	if v, e := ss.PopAt(1); e != EmptyErr {
		t.Errorf("expect err %v, but got %v\n", EmptyErr, e)
	} else if v != 0 {
		t.Errorf("expect v %v, but got %v\n", 0, v)
	}
	if v, e := ss.Pop(); e != EmptyErr {
		t.Errorf("expect err %v, but got %v\n", EmptyErr, e)
	} else if v != 0 {
		t.Errorf("expect v %v, but got %v\n", 0, v)
	}
}
