package ctci

import (
	"container/list"
	"testing"
)

func TestLastNth(t *testing.T) {
	l := list.New()
	cases := []struct {
		list   []int
		n      int
		expect interface{}
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4},
		{[]int{1, 2, 3, 4, 5}, 5, 1},
		{[]int{1, 2, 3}, 4, nil},
	}

	for _, c := range cases {
		l.Init()
		for _, e := range c.list {
			l.PushBack(e)
		}
		if r := LastNth(l, c.n); r != c.expect {
			t.Errorf("list:%v, n:%d, expect %v, but got %v\n",
				c.list, c.n, c.expect, r)
		}
	}
}
