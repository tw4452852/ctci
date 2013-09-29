package ctci

import (
	"fmt"
)

// Given an image represented by an NxN matrix,
// where each pixel in the image is 4 bytes,
// write a method to rotate the image by 90 degrees. Can you do this in place?

type matrix [][]int

func (m matrix) String() string {
	s := "["
	for i := 0; i < len(m); i++ {
		s += "["
		for j := 0; j < len(m); j++ {
			s += fmt.Sprintf("%d, ", m[i][j])
		}
		s += "]"
	}
	s += "]"
	return s
}

func (m matrix) IsEqual(n matrix) bool {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[i][j] != n[i][j] {
				return false
			}
		}
	}
	return true
}

func (m matrix) Transpose() {
	n := len(m) // assume n*n matrix
	// exchange according to the diagnal
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			swap(&m[i][j], &m[j][i])
		}
	}
	// exchange up and down
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			swap(&m[i][j], &m[n-i-1][j])
		}
	}
}

func swap(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}
