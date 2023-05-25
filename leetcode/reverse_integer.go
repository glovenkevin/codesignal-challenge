package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/reverse-integer/description/
func main() {
	tc := []struct {
		x   int
		res int
	}{
		// {123, 321},
		// {120, 21},
		{-123, -321},
	}

	for _, t := range tc {
		res := reverse(t.x)

		if res != t.res {
			fmt.Println("test failed")
			continue
		}

		fmt.Println("test success")
	}
}

func reverse(x int) int {
	isMinus := 1
	if x < 0 {
		x *= -1
		isMinus = -1
	}

	s := []byte(strconv.Itoa(x))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	ret, _ := strconv.Atoi(string(s))

	maxInt := 1 << 31
	if ret > maxInt {
		return 0
	}

	return ret * isMinus
}
