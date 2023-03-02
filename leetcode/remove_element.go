package main

import "fmt"

// source: https://leetcode.com/problems/remove-element/description/
func main() {
	tc := []struct {
		in  []int
		val int
		res int
	}{
		{
			in:  []int{3, 2, 2, 3},
			val: 3,
			res: 2,
		},
		{
			in:  []int{3, 2, 2, 3},
			val: 2,
			res: 2,
		},
		{
			in:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			val: 2,
			res: 5,
		},
		{
			in:  []int{1, 2, 3, 4},
			val: 4,
			res: 3,
		},
	}

	for _, t := range tc {
		res := removeElement(t.in, t.val)
		if res != t.res {
			fmt.Printf("Result: %d, Expected: %d \n", res, t.res)
			continue
		}
		fmt.Println("Test Success")
	}
}

func removeElement(nums []int, val int) int {
	var pos int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if nums[pos] == val {
				nums[pos], nums[i] = nums[i], nums[pos]
			}
			pos++
			continue
		}
	}

	return pos
}
