package main

import "fmt"

// source: https://leetcode.com/problems/search-insert-position/
func main() {
	tc := []struct {
		nums   []int
		target int
		res    int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			res:    2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			res:    1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			res:    4,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 0,
			res:    0,
		},
		{
			nums:   []int{1, 3},
			target: 2,
			res:    1,
		},
	}

	for _, t := range tc {
		res := searchInsert(t.nums, t.target)
		if res == t.res {
			fmt.Println("Test Success")
			continue
		}
		fmt.Printf("Result: %d, Expected: %d \n", res, t.res)
	}
}

// simplest solution is by using the binary tree method
func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums))
}

func binarySearch(nn []int, t, low, high int) int {
	if low >= high {
		if low != len(nn) && nn[low] < t {
			return low + 1
		}
		return low
	}

	mid := (low + high) / 2
	if nn[mid] > t {
		return binarySearch(nn, t, low, mid-1)
	}
	if nn[mid] < t {
		return binarySearch(nn, t, mid+1, high)
	}

	return mid
}
