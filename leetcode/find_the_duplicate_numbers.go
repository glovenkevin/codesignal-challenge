package main

// source: https://leetcode.com/problems/find-the-duplicate-number/
func main() {
}

func findDuplicate(nums []int) int {
	nn := make([]int, len(nums))
	for _, n := range nums {
		nn[n]++
		if nn[n] > 1 {
			return n
		}
	}
	return 0
}
