package ctci

type node struct {
	v    byte
	next *node
}

// Implement an algorithm to delete a node in the middle of a single linked list,
// given only access to that node.
//
// EXAMPLE
//
// Input: the node 'c' from the linked list a->b->c->d->e
// Result: nothing is returned, but the new linked list looks like a->b->d->e

func RemoveElement(de *node) {
	if de == nil {
		return
	}
	ne := de.next
	// tail element
	if ne == nil {
		de.v = 0
		return
	}
	de.next = ne.next
	de.v = ne.v
}

func ListToBytes(h *node) []byte {
	bs := make([]byte, 0)
	for h != nil {
		if h.v != 0 {
			bs = append(bs, h.v)
		}
		h = h.next
	}
	return bs
}

func BytesToList(bs []byte) *node {
	var head, tail *node
	for i, b := range bs {
		n := &node{b, nil}
		if i == 0 {
			head = n
			tail = n
			continue
		}
		tail.next = n
		tail = n
	}
	return head
}
