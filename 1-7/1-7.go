package ctci

import (
	"fmt"
)

// Write an algorithm such that if an element in an MxN matrix is 0,
// its entire row and column is set to 0.

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
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] != n[i][j] {
				return false
			}
		}
	}
	return true
}

func (m matrix) Zero() {
	rows := make([]bool, len(m))
	cols := make([]bool, len(m[0]))
	// check the zero rows and columns
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	// clear the rows and columns
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if rows[i] || cols[j] {
				m[i][j] = 0
			}
		}
	}
}

func swap(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}
