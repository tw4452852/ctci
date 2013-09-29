package ctci

import (
	"bytes"
	"testing"
)

func TestRemoveDuplicateElement(t *testing.T) {
	cases := []struct {
		input  []byte
		expect []byte
	}{
		{[]byte{1, 2, 3, 4, 2, 3, 6, 2, 1, 11}, []byte{1, 2, 3, 4, 6, 11}},
		{[]byte{1, 2, 3}, []byte{1, 2, 3}},
		{[]byte{1, 1, 2, 2, 2, 3, 3, 3, 3}, []byte{1, 2, 3}},
		{[]byte{}, []byte{}},
	}

	for _, c := range cases {
		if r := ListToBytes(RemoveDuplicateElement(BytesToList(c.input))); !bytes.Equal(r, c.expect) {
			t.Errorf("expect %v, but got %v\n", c.expect, r)
		}
	}
}
