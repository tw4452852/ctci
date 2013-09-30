package ctci

import (
	"bytes"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		input  []byte
		nth    int
		expect []byte
	}{
		{[]byte{1, 2, 3, 4}, 1, []byte{2, 3, 4}},
		{[]byte{1, 2, 3, 4}, 2, []byte{1, 3, 4}},
		{[]byte{1, 2, 3, 4}, 4, []byte{1, 2, 3}},
		{[]byte{1, 2, 3, 4}, 0, []byte{1, 2, 3, 4}},
	}

	for _, c := range cases {
		var de *node
		l := BytesToList(c.input)
		for i := 0; i < c.nth; i++ {
			if i == 0 {
				de = l
				continue
			}
			de = de.next
		}
		RemoveElement(de)
		if r := ListToBytes(l); !bytes.Equal(c.expect, r) {
			t.Errorf("input %v, delete %dth, expect %v, but got %v\n",
				c.input, c.nth, c.expect, r)
		}
	}
}
