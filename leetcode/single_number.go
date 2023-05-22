package main

import "fmt"

// source: https://leetcode.com/problems/single-number/
func main() {
	tc := []struct {
		nums []int
		res  int
	}{
		{
			nums: []int{2, 2, 1},
			res:  1,
		},
		{
			nums: []int{4, 1, 2, 1, 2},
			res:  4,
		},
	}

	for _, t := range tc {
		res := singleNumber(t.nums)

		if res != t.res {
			fmt.Println("test failed")
			continue
		}

		fmt.Println("test success")
	}
}

// note: do it with linear time & constant / O(1) space complexity
// it is trivial if you understand the bitwise operator
func singleNumber(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = res ^ nums[i]
	}

	return res
}
