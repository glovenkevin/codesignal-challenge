package main

import (
	"fmt"
	"regexp"
	"strings"
)

// source: https://leetcode.com/problems/valid-palindrome/
func main() {
	tc := []struct {
		s   string
		res bool
	}{
		{" ", true},
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for _, t := range tc {
		res := isPalindrome(t.s)

		if res != t.res {
			fmt.Println("test failed")
			continue
		}

		fmt.Println("test success")
	}
}

func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	s = strings.ToLower(s)
	r := regexp.MustCompile("[^a-z0-9]+|[\\s+]")
	bb := r.ReplaceAllLiteralString(s, "")
	s = string(bb)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
