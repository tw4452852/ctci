package ctci

// Implement a function to check if a tree is balanced.
// For the purposes of this question, a balanced tree is defined to be a tree
// such that no two leaf nodes differ in distance from the root by more than one.

type node struct {
	v           interface{}
	left, right *node
}

func IsBanlanceTree(root *node) bool {
	max, min := 0, 0
	getDepth(root, &max, &min, 0)
	return max-min <= 1
}

func getDepth(n *node, max, min *int, cur int) {
	if n == nil {
		return
	}
	if n.left == nil && n.right == nil {
		if cur > *max {
			*max = cur
		}
		if *min == 0 {
			*min = cur
		} else if cur < *min {
			*min = cur
		}
	}
	getDepth(n.left, max, min, cur+1)
	getDepth(n.right, max, min, cur+1)
}
