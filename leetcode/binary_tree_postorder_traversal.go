package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// source: https://leetcode.com/problems/binary-tree-postorder-traversal/
func main() {
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	getValPostOrder(root, &res)
	return res
}

func getValPostOrder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		getValPostOrder(node.Left, res)
	}

	if node.Right != nil {
		getValPostOrder(node.Right, res)
	}

	*res = append(*res, node.Val)
}
