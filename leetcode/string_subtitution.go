package main

import "fmt"

func main() {
	fmt.Println(isValid("abccba"))       // Expected: false
	fmt.Println(isValid("abcabcababcc")) // Expected: true
}

func isValid(s string) bool {

	if s == "" {
		return true
	}

	sizeS := len(s)
	for i := 0; i <= sizeS-3; i++ {
		tempChar := s[i : i+3]
		if tempChar == "abc" {
			fmt.Println(s[0:i] + s[i+3:])
			return isValid(s[0:i] + s[i+3:])
		}
	}

	return false
}
