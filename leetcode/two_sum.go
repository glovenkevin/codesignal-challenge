package main

import "fmt"

// url: https://leetcode.com/problems/two-sum/
// The other way to solve this, is that by sorting the array, than start the checking from the head and tail
func main() {
	pp := []struct {
		nums   []int
		target int
	}{
		{ // Expected: [0,1]
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		{ // Expected: [1,2]
			nums:   []int{3, 2, 4},
			target: 6,
		},
		{ // Expected: [0,1]
			nums:   []int{3, 3},
			target: 6,
		},
	}

	for _, p := range pp {
		fmt.Println(twoSum(p.nums, p.target))
	}
}

// Complexity: O(n*n)
func twoSum(nums []int, target int) []int {
	var ret []int
	for i, n := range nums {
		for k := i + 1; k < len(nums); k++ { // worst case: (n-1) == n
			if n+nums[k] == target {
				return []int{i, k}
			}
		}
	}

	return ret
}

// Complexity: O(n)
func twoSumQuickes(nums []int, target int) []int {
	index := make(map[int]int)
	for idx, num := range nums {
		i, ok := index[target-num]
		if !ok {
			index[num] = idx
			continue
		}

		return []int{i, idx}
	}

	return nil
}
