package ctci

import (
	"container/list"
)

// Write code to remove duplicates from an unsorted linked list.
//
// FOLLOW UP
//
// How would you solve this problem if a temporary buffer is not allowed?

func RemoveDuplicateElement(l *list.List) *list.List {
	clean := make([]*list.Element, 0)
	for p := l.Front(); p != nil; p = p.Next() {
		for q := p.Next(); q != nil; q = q.Next() {
			if p.Value == q.Value {
				clean = append(clean, q)
			}
		}
	}
	for _, e := range clean {
		l.Remove(e)
	}
	return l
}

func ListToBytes(l *list.List) []byte {
	bs := make([]byte, 0, l.Len())
	for p := l.Front(); p != nil; p = p.Next() {
		bs = append(bs, p.Value.(byte))
	}
	return bs
}

func BytesToList(bs []byte) *list.List {
	l := list.New()
	for _, b := range bs {
		l.PushBack(b)
	}
	return l
}
