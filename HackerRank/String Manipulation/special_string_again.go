package main

import (
	"fmt"
)

func main() {
	// fmt.Println(substrCount(5, "asasd"))   // Expected: 7
	// fmt.Println(substrCount(7, "abcbaba")) // Expected: 10
	fmt.Println(substrCount(4, "aaaa")) // Expected: 10
}

/*
	RULE SPECIAL STRING:
		- All of the characters are the same, e.g. aaa.
		- All characters except the middle one are the same, e.g. aadaa.
*/
func substrCount(n int32, s string) int64 {

	var (
		count    int64  = int64(n)
		word     string = ""
		idx      int32  = 2
		startIdx int32  = 0
	)

	if isPalindrome(s) {
		count++
	}

	for idx < n {

		if idx+startIdx <= n {

			word = s[startIdx : idx+startIdx]
			if isPalindrome(word) {
				count++
			}
			startIdx++

		} else {
			idx++
			startIdx = 0
		}

	}

	return count
}

func isPalindrome(s string) bool {
	return s == reverseSpecial(s)
}

func reverseSpecial(s string) string {
	strLen := len([]rune(s))
	half := strLen / 2
	newWord := ""
	if strLen%2 == 1 {
		half = (strLen + 1) / 2
		newWord = s[half:] + s[half-1:half] + reverse(s[:half-1])
	} else {
		newWord = s[half:] + reverse(s[:half])
	}

	return newWord
}

func reverse(s string) string {
	rs := []rune(s)
	lenStr := len(rs)
	for i, j := 0, lenStr-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
