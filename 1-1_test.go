package ctci

import (
	"testing"
)

func TestIsUnique(t *testing.T) {
	cases := []struct {
		input  string
		expect bool
	}{
		{"abcd", true},
		{"", true},
		{`_\/,"'.;`, true},
		{"abca", false},
		{"aaaa", false},
	}

	for _, c := range cases {
		r := IsUnique(c.input)
		if r != c.expect {
			t.Errorf("input %q, expect %v, but get %v\n",
				c.input, c.expect, r)
		}
	}
}
