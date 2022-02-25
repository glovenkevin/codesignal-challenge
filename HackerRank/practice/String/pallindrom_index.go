package main

import "fmt"

func main() {
	fmt.Println(palindromeIndex("bcbc")) // Expected: 0/3

	// fmt.Println(palindromeCheck("baab", 4))
}

func palindromeIndex(s string) int32 {

	sSize := len(s)
	if palindromeCheck(s, sSize) {
		return -1
	}

	for i, j := 0, sSize-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {

			newS := s[:i] + s[i+1:]
			if palindromeCheck(newS, len(newS)) {
				return int32(i)
			}

			newS = s[:j] + s[j+1:]
			if palindromeCheck(newS, len(newS)) {
				return int32(j)
			}

		}
	}

	return -1
}

func palindromeCheck(s string, sz int) bool {
	l := ""
	r := ""
	mid := sz / 2
	if sz%2 == 1 {
		l = s[:mid]
		r = s[mid+1:]
	} else {
		l = s[:mid]
		r = s[mid:]
	}

	for i := 0; i < len(l); i++ {
		if r[i] != l[mid-1-i] {
			return false
		}
	}

	return true
}
