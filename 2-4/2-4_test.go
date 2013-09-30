package ctci

import (
	"bytes"
	"testing"
)

func TestListAdd(t *testing.T) {
	cases := []struct {
		inputs [][]byte
		expect []byte
	}{
		{
			[][]byte{
				[]byte{5, 1, 4},
				[]byte{5, 8, 5},
			},
			[]byte{0, 0, 0, 1},
		},
		{
			[][]byte{
				[]byte{},
				[]byte{5, 8, 5},
			},
			[]byte{5, 8, 5},
		},
		{
			[][]byte{
				[]byte{},
				[]byte{},
			},
			[]byte{},
		},
		{
			[][]byte{
				[]byte{1, 2},
				[]byte{2, 8, 2},
			},
			[]byte{3, 0, 3},
		},
	}

	for _, c := range cases {
		rl := ListAdd(BytesToList(c.inputs[0]), BytesToList(c.inputs[1]))
		if r := ListToBytes(rl); !bytes.Equal(r, c.expect) {
			t.Errorf("input %v, %v, expect %v, but got %v\n",
				c.inputs[0], c.inputs[1], c.expect, r)
		}
	}
}
