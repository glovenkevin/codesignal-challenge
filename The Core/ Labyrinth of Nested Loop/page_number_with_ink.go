package main

import "fmt"

func main() {
	fmt.Println(pagesNumberingWithInk(1, 5))  // expected: 5
	fmt.Println(pagesNumberingWithInk(21, 5)) // expected: 22
	fmt.Println(pagesNumberingWithInk(8, 4))  // expected: 11
}

func pagesNumberingWithInk(current int, numberOfDigits int) int {
	numberLength := len(fmt.Sprint(current))
	for numberOfDigits >= numberLength {
		numberOfDigits -= numberLength
		current++
		numberLength = len(fmt.Sprint(current))
	}
	return current - 1
}
