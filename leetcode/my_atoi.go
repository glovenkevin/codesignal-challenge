package main

import (
	"fmt"
	"strings"
)

// source: https://leetcode.com/problems/string-to-integer-atoi/
func main() {
	tc := []struct {
		s   string
		res int
	}{
		{"42", 42},
		{"    -42", -42},
		{"words and 987", 0},
		{"4193 with words", 4193},
		{"-91283472332", -2147483648},
		{"+1", 1},
	}

	for _, t := range tc {
		res := myAtoi(t.s)

		if res != t.res {
			fmt.Println("test failed")
			continue
		}

		fmt.Println("test success")
	}
}

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	var (
		multiply = 10
		res      = 0
		sign     = 1
		maxInt   = 1 << 31
	)
	for i, r := range s {
		if r == '-' && i == 0 {
			sign *= -1
			continue
		}
		if r == '+' && i == 0 {
			continue
		}
		if r > '9' || r < '0' {
			break
		}

		res = (res * multiply) + int(r-'0')
		if res >= maxInt {
			res = maxInt
			if sign < 0 {
				res--
			}
			break
		}
	}

	return res * sign
}
