package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/length-of-last-word/
func main() {
	tc := []struct {
		s string
		l int
	}{
		{"Hellow World", 5},
		{"   fly me   to   the moon  ", 4},
		{"a", 1},
	}
	for _, c := range tc {
		res := lengthOfLastWord(c.s)
		if res != c.l {
			fmt.Printf("Result: %d, expected: %d\n", res, c.l)
		} else {
			fmt.Println("Test successful")
		}
	}
}

func lengthOfLastWord(s string) int {
	var l int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && l == 0 {
			continue
		}
		if s[i] == ' ' && l > 0 {
			return l
		}

		l++
	}

	return l
}
