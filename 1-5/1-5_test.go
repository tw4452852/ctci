package ctci

import (
	"bytes"
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	cases := []struct {
		input  []byte
		expect []byte
	}{
		{[]byte("a b c"), []byte("a%20b%20c")},
		{[]byte(""), []byte("")},
		{[]byte("a  c"), []byte("a%20%20c")},
	}

	for _, c := range cases {
		r := ReplaceSpace(c.input)
		if !bytes.Equal(r, c.expect) {
			t.Errorf("input %v, expect %v, but got %v\n",
				c.input, c.expect, r)
		}
	}
}
