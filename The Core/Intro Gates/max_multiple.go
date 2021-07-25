package main

import "fmt"

func main() {
	fmt.Println(maxMultiple(3, 10)) // expected: 9, because 9 possible divided by 3 and 9 != larger than 10 and 9 is not zero
}

func maxMultiple(divisor int, bound int) int {
	return bound - (bound % divisor)
}
