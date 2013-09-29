package ctci

import (
	"container/list"
)

// Implement an algorithm to find the nth to last element of a singly linked list.

func LastNth(l *list.List, n int) interface{} {
	if n > l.Len() {
		return nil
	}
	f := l.Front()
	b := f
	for i := 0; i < n; i++ {
		b = b.Next()
	}
	for b != nil {
		f = f.Next()
		b = b.Next()
	}
	return f.Value
}
