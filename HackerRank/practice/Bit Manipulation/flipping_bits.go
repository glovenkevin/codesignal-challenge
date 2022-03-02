package main

import (
	"fmt"
)

func main() {
	fmt.Println(flippingBits(4)) // Expected: 4294967291
}

func flippingBits(n int64) int64 {
	return int64(uint32(^n))
}
