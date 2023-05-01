package main

import "fmt"

// source: https://leetcode.com/problems/binary-tree-inorder-traversal
func main() {
	tc := []struct {
		root *TreeNode
		res  []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			res: []int{1, 3, 2},
		},
		{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			res: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}

	for _, t := range tc {
		res := inorderTraversal(t.root)

		if len(t.res) != len(res) {
			fmt.Println("test failed")
			continue
		}

		for i, n := range res {
			if t.res[i] != n {
				fmt.Println("test failed")
			}
		}

		fmt.Println("test success")
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	diveInorderTraversal(root, &res)
	return res
}

func diveInorderTraversal(root *TreeNode, ii *[]int) {
	if root.Left != nil {
		diveInorderTraversal(root.Left, ii)
	}

	*ii = append(*ii, root.Val)

	if root.Right != nil {
		diveInorderTraversal(root.Right, ii)
	}
}
