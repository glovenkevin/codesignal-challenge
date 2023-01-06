package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/valid-parentheses/
func main() {
	pp := []struct {
		s     string
		valid bool
	}{
		{s: "[]{}()", valid: true},
		{s: "()", valid: true},
		{s: "(]", valid: false},
		{s: "{[]}", valid: true},
		{s: "(])", valid: false},
		{s: "([}}])", valid: false},
	}
	for i, p := range pp {
		fmt.Printf("%d. Result: %v -> Actual: %v \n", i+1, isValid(p.s), p.valid)
	}
}

func isValid(s string) bool {
	parenthesesCouple := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	var stack []string
	for _, parenthese := range s {
		ps := string(parenthese)
		if len(stack) == 0 {
			stack = append(stack, ps)
			continue
		}

		c, ok := parenthesesCouple[ps]
		if !ok {
			stack = append(stack, ps)
			continue
		}
		if ok && stack[len(stack)-1] == c {
			stack = stack[:len(stack)-1]
			continue
		}
		if ok && stack[len(stack)-1] != c {
			stack = append(stack, ps)
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
