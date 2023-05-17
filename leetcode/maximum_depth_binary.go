package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// source: https://leetcode.com/problems/maximum-depth-of-binary-tree/
func main() {
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res, h int = 1, 1
	checkDepth(root, &res, &h)
	return res
}

func checkDepth(n *TreeNode, total *int, h *int) {
	*total++
	if n.Left != nil {
		checkDepth(n.Left, total, h)
	}

	if n.Right != nil {
		checkDepth(n.Right, total, h)
	}

	if *total > *h {
		*h = *total
	}
	*total--
}
