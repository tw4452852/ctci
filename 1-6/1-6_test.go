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
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
				[]int{13, 14, 15, 16},
			},
			matrix{
				[]int{4, 8, 12, 16},
				[]int{3, 7, 11, 15},
				[]int{2, 6, 10, 14},
				[]int{1, 5, 9, 13},
			},
		},
	}

	for _, c := range cases {
		c.input.Transpose()
		if !c.expect.IsEqual(c.input) {
			t.Errorf("expect %v, but got %v\n",
				c.expect, c.input)
		}
	}
}
