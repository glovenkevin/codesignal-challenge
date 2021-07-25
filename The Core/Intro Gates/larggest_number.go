package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(2) // expected: 99
}

func largestNumber(n int) int {
	return int(math.Pow10(n)) - 1
}
