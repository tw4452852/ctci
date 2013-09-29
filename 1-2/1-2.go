package ctci

// Write code to reverse a C-Style String.
// (C-String means that "bcd" represented as five characters, including the null character.)

func RevertCString(cs []byte) {
	for i := 0; i < (len(cs)-1)/2; i++ {
		swap(&cs[i], &cs[len(cs)-2-i])
	}
}

func swap(a, b *byte) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
