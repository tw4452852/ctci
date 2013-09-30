package ctci

import (
	"testing"
)

func TestFindCycleBegin(t *testing.T) {
	cases := []struct {
		list       []byte
		beginIndex int
	}{
		{[]byte{1, 2, 3, 4}, -1},
		{[]byte{1, 2, 3, 4}, 0},
		{[]byte{1, 2, 3, 4}, 2},
		{[]byte{}, -1},
	}

	for _, c := range cases {
		head := BytesToList(c.list)
		var expect *node
		if c.beginIndex >= 0 {
			for i, p := 0, head; p != nil; i, p = i+1, p.next {
				if i == c.beginIndex {
					expect = p
				}
				if p.next == nil {
					p.next = expect
					break
				}
			}
		}
		if r := FindCycleBegin(head); r != expect {
			t.Errorf("expect %v, but got %v\n", expect, r)
		}
	}
}

func BytesToList(bs []byte) *node {
	var head, tail *node
	for i, b := range bs {
		n := &node{b, nil}
		if i == 0 {
			head = n
			tail = n
			continue
		}
		tail.next = n
		tail = n
	}
	return head
}
