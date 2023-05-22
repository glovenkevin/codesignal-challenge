package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// source: https://leetcode.com/problems/maximum-depth-of-binary-tree/
func main() {
	tc := []struct {
		root *TreeNode
		res  int
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			},
			res: 3,
		},
	}

	for _, t := range tc {
		res := maxDepth(t.root)

		if res == t.res {
			fmt.Println("test success")
			continue
		}

		fmt.Println("test failed")
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res, high int
	checkDepth(root, &res, &high)
	return high
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
