package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(deleteDigit(152))  // expected: 52
	fmt.Println(deleteDigit(1001)) // expected: 101
}

func deleteDigit(n int) int {
	strItem := strconv.Itoa(n)
	arrString := []rune(strItem)
	biggerValue := 0
	for idx, _ := range arrString {
		newStr := ""
		for i, val := range arrString {
			if i != idx {
				newStr += string(val)
			}
		}
		newInt, _ := strconv.Atoi(newStr)
		if newInt > biggerValue {
			biggerValue = newInt
		}
	}
	return biggerValue
}
