package main

import "fmt"

// source: https://leetcode.com/problems/plus-one/
func main() {
	tc := []struct {
		dg  []int
		ret []int
	}{
		{
			dg:  []int{1, 2, 3},
			ret: []int{1, 2, 4},
		},
		{
			dg:  []int{9},
			ret: []int{1, 0},
		},
		{
			dg:  []int{2},
			ret: []int{3},
		},
		{
			dg:  []int{9, 9},
			ret: []int{1, 0, 0},
		},
	}

	for _, c := range tc {
		xx := plusOne(c.dg)

		if len(xx) != len(c.ret) {
			fmt.Printf("lenght is not equal: %v | expected: %v\n", xx, c.ret)
			continue
		}

		for i, x := range xx {
			if c.ret[i] != x {
				fmt.Printf("[FAIL] res: %v | expected: %v\n", xx, c.ret)
				break
			}
		}

		fmt.Println("test successful")
	}
}

func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1] += 1
	if digits[n-1] <= 9 {
		return digits
	}

	for i := n - 1; i > 0; i-- {
		if digits[i] == 10 {
			digits[i] = 0
			digits[i-1] += 1
		}
	}

	if digits[0] == 10 {
		digits[0] = 0
		var ret = make([]int, n+1)
		ret[0] = 1
		copy(ret[1:], digits)
		return ret
	}

	return digits
}

func plusOneSimpler(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		}

		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
