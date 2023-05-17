package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/longest-substring-without-repeating-characters/
func main() {
	tc := []struct {
		s   string
		res int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{" ", 1},
		{"au", 2},
		{"aab", 2},
		{"dvdf", 3},
		{"fbojelwjgercer", 7},
	}

	for _, t := range tc {
		res := lengthOfLongestSubstring(t.s)

		if t.res != res {
			fmt.Println("test failed")
			continue
		}

		fmt.Println("test success")
	}
}

// improvement using sliding windows
func lengthOfLongestSubstring(s string) int {
	var (
		runeMap          = make(map[rune]struct{}, len(s))
		length, first, i int
	)
	for i = 0; i < len(s); i++ {
		// log.Println(string(s[i]))
		r := rune(s[i])
		_, ok := runeMap[r]
		if ok {
			if length < len(s[first:i+1]) {
				length = len(s[first:i])
			}

			for rune(s[first]) != r {
				delete(runeMap, rune(s[first]))
				first++
			}
			first++

			// log.Println("reset")
			// log.Println("length", length)
			continue
		}

		runeMap[r] = struct{}{}
	}

	if len(s[first:i]) > length {
		length = len(s[first:i])
	}

	return length
}

// failed: time take to execute to long but success
func lengthOfLongestSubstringReal(s string) int {
	var (
		runeMap       = make(map[rune]int, len(s))
		length, first int
	)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		idx, ok := runeMap[r]
		if ok {
			if length < len(s[first:i+1]) {
				length = len(s[first:i])
			}
			runeMap = make(map[rune]int, len(s))
			first = idx + 1
			i = idx
			continue
		}

		runeMap[r] = i
	}

	if len(runeMap) > length {
		length = len(runeMap)
	}

	return length
}

// using sliding window
func findLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	window := make([]bool, 128)

	var begin, end, low int

	for high := 0; high < len(s); high++ {
		if window[s[high]] {
			for s[low] != s[high] {
				window[s[low]] = false
				low++
			}

			low++
		} else {
			window[s[high]] = true

			if end-begin < high-low {
				begin = low
				end = high
			}
		}
	}

	return len(s[begin : end+1])
}
