package main

// soure: https://leetcode.com/problems/majority-element/
func main() {
}

// target: time O(n) && space O(1)
// target: time O(n) && space O(1)
func majorityElement(nums []int) int {
	var count, el int
	for _, n := range nums {
		if el == n {
			count++
			continue
		}

		if count == 0 {
			el = n
			count++
		} else {
			count--
		}
	}

	return el
}
