package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPower(125)) // expected: true
	fmt.Println(isPower(72))  // expected: false
}

// Source: https://stackoverflow.com/a/4429063/6759373
func isPower(n int) bool {
	if n == 1 {
		return true
	}

	a := 2
	for {
		b := n
		for b%a == 0 {
			b /= a
		}

		if a == n {
			return false
		} else if b == 1 {
			return true
		}
		a++
	}
}
