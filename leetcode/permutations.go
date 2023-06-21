package main

import "fmt"

// source: https://leetcode.com/problems/permutations/
func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var res [][]int
	findPermutation(nums, 0, len(nums)-1, &res)
	return res
}

func findPermutation(nn []int, left, right int, res *[][]int) {
	if left == right {
		nns := make([]int, len(nn))
		copy(nns, nn)
		*res = append(*res, nns)
		return
	}

	for i := left; i <= right; i++ {
		nn[left], nn[i] = nn[i], nn[left]

		findPermutation(nn, left+1, right, res)

		nn[left], nn[i] = nn[i], nn[left]
	}
}
