package ctci

// You have two numbers represented by a linked list,
// where each node contains a single digit.
// The digits are stored in reverse order,
// such that the 1's digit is at the head of the list.
// Write a function that adds the two numbers and returns the sum as a linked list.
//
// EXAMPLE
//
// Input: (3 -> 1 -> 5), (5 -> 9 -> 2)
//
// Output: 8 -> 0 -> 8

type node struct {
	v    byte
	next *node
}

func ListAdd(a, b *node) *node {
	var carry byte
	bs := []byte{}
	for a != nil || b != nil || carry != 0 {
		sum := carry
		if a != nil {
			sum += a.v
			a = a.next
		}
		if b != nil {
			sum += b.v
			b = b.next
		}
		carry = 0
		if sum >= 10 {
			sum %= 10
			carry = 1
		}
		bs = append(bs, sum)
	}
	return BytesToList(bs)
}

func ListToBytes(h *node) []byte {
	bs := make([]byte, 0)
	for h != nil {
		bs = append(bs, h.v)
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
