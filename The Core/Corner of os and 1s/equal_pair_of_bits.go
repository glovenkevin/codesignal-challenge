package main

import (
	"fmt"
)

func main() {
	fmt.Println(equalPairOfBits(1073741824, 1006895103))
}

func equalPairOfBits(n int, m int) int {
	return (^(n ^ m)) & ((n ^ m) + 1)
}
