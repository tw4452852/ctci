package ctci

import (
	"testing"
)

func TestIsRotation(t *testing.T) {
	cases := []struct {
		inputs [2]string
		expect bool
	}{
		{[2]string{"waterbottler", "erbottlerwat"}, true},
		{[2]string{"waterbottler", "erbottlerwa"}, false},
		{[2]string{"", ""}, true},
	}

	for _, c := range cases {
		if r := IsRotation(c.inputs[0], c.inputs[1]); r != c.expect {
			t.Errorf("input[2]: %q, %q, expect %v, but got %v\n",
				c.inputs[0], c.inputs[1], c.expect, r)
		}
	}
}
