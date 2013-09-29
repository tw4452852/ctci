package ctci

import (
	"bytes"
	"testing"
)

func TestRevertCString(t *testing.T) {
	cases := []struct {
		cs     []byte
		expect []byte
	}{
		{[]byte{'a', 'b', 'c', 0}, []byte{'c', 'b', 'a', 0}},
		{[]byte{'a', 'c', 0}, []byte{'c', 'a', 0}},
		{[]byte{0}, []byte{0}},
	}

	for _, c := range cases {
		RevertCString(c.cs)
		if !bytes.Equal(c.expect, c.cs) {
			t.Errorf("expect %v, but got %v\n",
				c.expect, c.cs)
		}
	}
}
