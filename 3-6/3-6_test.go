package ctci

import (
	"testing"
)

func TestSortStack(t *testing.T) {
	cases := []struct {
		inputs []int
		expect []int
	}{
		{[]int{2, 6, 3, 5, 4}, []int{2, 3, 4, 5, 6}},
	}
	for _, c := range cases {
		stk := NewStack()
		for _, ci := range c.inputs {
			stk.Push(ci)
		}
		rs := SortStack(stk)
		for i := len(c.expect) - 1; i >= 0; i-- {
			if r := rs.Pop(); c.expect[i] != r {
				t.Errorf("expect %v, but got %v\n", c.expect[i], r)
			}
		}
	}
}
