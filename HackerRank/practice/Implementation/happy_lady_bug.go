package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := []string{
		"AABBC_",     // Expected: "NO"
		"AABBC_C",    // Expected: "YES"
		"DD__FQ_QQF", // Expected: "YES"
		"_",          // Expected: "YES"
		"AABCBC",     // Expected: "NO"
		"B_RRBR",     // Expected: "YES"
		"X_Y__X",     // Expected: "NO"
		"AABBCC",     // Expected: "YES"
	}

	for _, s := range ss {
		fmt.Println(happyLadybugs(s, int32(len(s))))
	}
}

func happyLadybugs(bb string, n int32) string {
	var emptyCell = strings.Count(bb, "_")
	for i, b := range bb {
		w := string(b)
		if w == "_" || (i != 0 && string(bb[i-1]) == w) || (i != len(bb)-1 && string(bb[i+1]) == w) {
			continue
		}

		total := strings.Count(bb, w)
		if total < 2 {
			return "NO"
		}

		if emptyCell == 0 && ((i != 0 && string(bb[i-1]) != w) || (i != len(bb)-1 && string(bb[i+1]) != w)) {
			return "NO"
		}
	}

	return "YES"
}
