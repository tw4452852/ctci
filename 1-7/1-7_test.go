package ctci

import (
	"testing"
)

func TestTranspose(t *testing.T) {
	cases := []struct {
		input  matrix
		expect matrix
	}{
		{
			matrix{
				[]int{1, 2, 3, 4, 5},
				[]int{5, 0, 7, 8, 0},
				[]int{9, 0, 0, 2, 6},
				[]int{3, 4, 5, 1, 2},
			},
			matrix{
				[]int{1, 0, 0, 4, 0},
				[]int{0, 0, 0, 0, 0},
				[]int{0, 0, 0, 0, 0},
				[]int{3, 0, 0, 1, 0},
			},
		},
	}

	for _, c := range cases {
		c.input.Zero()
		if !c.expect.IsEqual(c.input) {
			t.Errorf("expect %v, but got %v\n",
				c.expect, c.input)
		}
	}
}
