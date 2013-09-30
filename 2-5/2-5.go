package ctci

// Given a circular linked list,
// implement an algorithm which returns node at the beginning of the loop.
//
// DEFINITION
//
// Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node,
// so as to make a loop in the linked list.
//
// EXAMPLE
//
// Input: A -> B -> C -> D -> E -> C [the same C as earlier]
//
// Output: C

type node struct {
	v    byte
	next *node
}

func FindCycleBegin(head *node) *node {
	fast, slow := head, head
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return slow
}
