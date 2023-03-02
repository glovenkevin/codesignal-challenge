package main

import "fmt"

// source: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func main() {
	pp := []struct {
		in       []int
		expected []int
		k        int
	}{
		{in: []int{1, 1, 2}, expected: []int{1, 2}, k: 2},
		{in: []int{1}, expected: []int{1}, k: 1},
		{in: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expected: []int{0, 1, 2, 3, 4}, k: 5},
		{in: []int{1, 2}, expected: []int{1, 2}, k: 2},
		{in: []int{1, 1}, expected: []int{1}, k: 1},
		{in: []int{1, 2, 3}, expected: []int{1, 2, 3}, k: 3},
		{in: []int{1, 2, 2}, expected: []int{1, 2}, k: 2},
	}
	for _, p := range pp {
		fmt.Printf("Result k: %d -> Actual k: %d \n", removeDuplicatesSimpler(p.in), p.k)
		fmt.Printf("Arr ex: %v -> Arr out: %v \n", p.expected, p.in)
		fmt.Println()
	}
}

func removeDuplicates(nums []int) int {
	var startIndexSearch int = 1
	for i := 0; i < len(nums)-1; i++ {
		var found bool
		for j := startIndexSearch; j < len(nums); j++ {
			if nums[i] < nums[j] && i+1 == j {
				found = true
				break
			}

			if nums[i] < nums[j] && i+1 != j {
				nums[i+1], nums[j] = nums[j], nums[i+1]
				found = true
				startIndexSearch = j + 1
				break
			}
		}

		if !found {
			return i + 1
		}
	}

	return len(nums)
}

func removeDuplicatesSimpler(nums []int) int {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
