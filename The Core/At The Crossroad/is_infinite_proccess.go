package main

import "fmt"

func main() {
	fmt.Println(isInfiniteProcess(2, 6)) // expected: true
}

func isInfiniteProcess(a int, b int) bool {
	if a > b {
		return true
	} else {
		return !((a-b)%2 == 0)
	}
}
