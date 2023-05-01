package main

import "fmt"

// source: https://leetcode.com/problems/same-tree/
func main() {
	tc := []struct {
		p   *TreeNode
		q   *TreeNode
		res bool
	}{
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			q: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			res: true,
		},
		{
			p: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			q: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			res: false,
		},
	}

	for _, t := range tc {
		res := isSameTree(t.p, t.q)

		if res != t.res {
			fmt.Println("test failed")
			continue
		}

		fmt.Println("test success")
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
