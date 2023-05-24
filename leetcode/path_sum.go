package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// source: https://leetcode.com/problems/path-sum/
func main() {
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	return getPathSum(root, 0, targetSum)
}

func getPathSum(root *TreeNode, sum int, targetSum int) bool {
	if root.Left != nil {
		res := getPathSum(root.Left, sum+root.Val, targetSum)
		if res {
			return res
		}
	}

	if root.Right != nil {
		res := getPathSum(root.Right, sum+root.Val, targetSum)
		if res {
			return res
		}
	}

	if root.Left == nil && root.Right == nil && sum+root.Val == targetSum {
		return true
	}

	return false
}

func hasPathSumSimplest(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
