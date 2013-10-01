package ctci

import (
	"testing"
)

func TestHanoi(t *testing.T) {
	if r, s := HanoiRecursion(6, 1, 2, 3), HanoiStack(6, 1, 2, 3); r != s {
		t.Errorf("recursion result %q != stack result %q\n", r, s)
	}
}
