package main

import "fmt"

// source: https://leetcode.com/problems/longest-common-prefix/
func main() {
	pp := []struct {
		ss []string
		r  string
	}{
		{ss: []string{"flower", "flow", "flight"}, r: "fl"},
		{ss: []string{"dog", "racecar", "car"}, r: ""},
		{ss: []string{"a"}, r: "a"},
		{ss: []string{"ab", "a"}, r: "a"},
	}
	for i, p := range pp {
		fmt.Printf("%d. Result: %s -> Actual: %s \n", i+1, longestCommonPrefix(p.ss), p.r)
	}
}

func longestCommonPrefix(strs []string) string {
	var cp string
	first := strs[0]
	for i, r := range first {
		var tmp string = string(r)
		for _, w := range strs[1:] {
			if i == len(w) || r != rune(w[i]) {
				tmp = ""
				break
			}
		}

		if tmp == "" {
			break
		}
		cp += tmp
	}

	return cp
}
