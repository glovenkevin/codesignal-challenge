package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(makeAnagram("cde", "dcf")) // expected: 2
	fmt.Println(makeAnagram("cde", "abc")) // expected: 4
}

func makeAnagram(a string, b string) int32 {

	aStr := a
	for _, rune := range []rune(aStr) {
		if strings.ContainsRune(b, rune) {
			a = strings.Replace(a, string(rune), "", 1)
			b = strings.Replace(b, string(rune), "", 1)
		}
	}

	sizeNoPair := len(a) + len(b)
	return int32(sizeNoPair)
}
