package ctci

import (
	"testing"
)

func TestIsBanlanceTree(t *testing.T) {
	nodes := []*node{}
	for i := 0; i < 10; i++ {
		nodes = append(nodes, &node{v: i})
	}
	// 0->1,2
	// 2->3
	nodes[0].left = nodes[1]
	nodes[0].right = nodes[2]
	nodes[2].right = nodes[3]
	if r := IsBanlanceTree(nodes[0]); r != true {
		t.Errorf("expect %v, but got %v\n", true, r)
	}
	// 3->4
	nodes[3].right = nodes[4]
	if r := IsBanlanceTree(nodes[0]); r != false {
		t.Errorf("expect %v, but got %v\n", false, r)
	}
	// 1->5
	nodes[1].left = nodes[5]
	if r := IsBanlanceTree(nodes[0]); r != true {
		t.Errorf("expect %v, but got %v\n", true, r)
	}
}
