package ctci

import (
	"testing"
)

func TestMyQueue(t *testing.T) {
	q := NewMyQueue()
	if v, e := q.DeQueue(); e != EmptyErr {
		t.Errorf("expect err %v, but got %v\n", EmptyErr, e)
	} else if v != nil {
		t.Errorf("expect value %v, but got %v\n", nil, v)
	}
	q.EnQueue(1)
	if v, e := q.DeQueue(); e != nil {
		t.Errorf("expect err %v, but got %v\n", nil, e)
	} else if v != 1 {
		t.Errorf("expect value %v, but got %v\n", 1, v)
	}
	q.EnQueue(2)
	q.EnQueue(3)
	if v, e := q.DeQueue(); e != nil {
		t.Errorf("expect err %v, but got %v\n", nil, e)
	} else if v != 2 {
		t.Errorf("expect value %v, but got %v\n", 2, v)
	}
	if v, e := q.DeQueue(); e != nil {
		t.Errorf("expect err %v, but got %v\n", nil, e)
	} else if v != 3 {
		t.Errorf("expect value %v, but got %v\n", 3, v)
	}
}
