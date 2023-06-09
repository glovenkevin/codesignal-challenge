package main

// source: https://leetcode.com/problems/contains-duplicate/
func main() {
}

// notes: other way to do it easily is by sort the number first, than just do a checking
func containsDuplicate(nums []int) bool {
	count := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := count[n]; ok {
			return true
		}

		count[n] = struct{}{}
	}

	return false
}
