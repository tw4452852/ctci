package ctci

import (
	"bytes"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	cases := []struct {
		cs     []byte
		expect []byte
	}{
		{[]byte{'a', 'b', 'c', 0}, []byte{'a', 'b', 'c', 0}},
		{[]byte{'a', 'a', 'a', 0}, []byte{'a', 0}},
		{[]byte{0}, []byte{0}},
		{[]byte{'a', 'a', 'b', 'b', 0}, []byte{'a', 'b', 0}},
		{[]byte{'a', 'b', 'a', 'b', 0}, []byte{'a', 'b', 0}},
	}

	for _, c := range cases {
		r := RemoveDuplicate(c.cs)
		if !bytes.Equal(r, c.expect) {
			t.Errorf("expect %v, but got %v\n", c.expect, r)
		}
	}
}
