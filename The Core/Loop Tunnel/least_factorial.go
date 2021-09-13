package main

import "fmt"

func main() {
	fmt.Println(leastFactorial(17)) // expected: 24
}

func leastFactorial(n int) int {
	var idx, sum int = 1, 1
	for {
		if sum < n {
			sum *= idx
			idx++
		} else {
			return sum
		}
	}
}
