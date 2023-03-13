package main

import "fmt"

// source: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func main() {
	tc := []struct {
		h, n string
		x    int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"a", "a", 0},
	}

	for _, c := range tc {
		x := strStr(c.h, c.n)
		if x != c.x {
			fmt.Printf("[FAILED] result: %v | expected: %v\n", x, c.x)
			continue
		}
		fmt.Println("test successful")
	}
}

func strStr(haystack string, needle string) int {
	for i, r := range []byte(haystack) {
		if r == needle[0] && i+len(needle) <= len(haystack) && haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
