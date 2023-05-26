package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// source: https://leetcode.com/problems/binary-tree-preorder-traversal/
func main() {
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	getValPreOrder(root, &res)
	return res
}

func getValPreOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)

	if node.Left != nil {
		getValPreOrder(node.Left, res)
	}

	if node.Right != nil {
		getValPreOrder(node.Right, res)
	}
}
