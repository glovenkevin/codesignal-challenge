package main

import (
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/palindrome-number/
func main() {
	pp := []int{
		121,  // true
		-121, // false
	}
	for _, p := range pp {
		fmt.Println(isPalindrome(p))
	}
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
