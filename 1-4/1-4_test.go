package ctci

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		as     []byte
		bs     []byte
		expect bool
	}{
		{[]byte("abcd"), []byte("abdd"), false},
		{[]byte("abbcd"), []byte("abcdb"), true},
		{[]byte(""), []byte(""), true},
	}

	for _, c := range cases {
		r := IsAnagram(c.as, c.bs)
		if r != c.expect {
			t.Errorf("input %q, %q, expect %v, but got %v\n",
				c.as, c.bs, c.expect, r)
		}
	}
}
