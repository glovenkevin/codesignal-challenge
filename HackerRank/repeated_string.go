package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedString("aba", 10)) // expected: 7
}

func repeatedString(s string, n int64) int64 {

	if !strings.Contains(s, "a") {
		return 0
	}

	var lnStr int = len(s)
	if lnStr == 1 {
		return n
	}

	repeat := int(n) / lnStr
	totalA := lnStr - len(strings.ReplaceAll(s, "a", ""))

	suffix := int(n) - (lnStr * repeat)
	sum := (totalA * repeat) + (suffix - len(strings.ReplaceAll(s[:suffix], "a", "")))
	return int64(sum)

}
