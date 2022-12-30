package main

import "fmt"

func main() {
	pp := []int64{
		4,  // expected: 6
		9,  // expected: 1
		21, // expected: 1
		12, // expected: 10
	}

	for _, p := range pp {
		fmt.Println(strangeCounter(p))
	}
}

func strangeCounter(t int64) int64 { // Timeout
	base := int64(3)
	for {
		if t-base <= 0 {
			break
		}

		t -= base
		base *= 2
	}

	return base - t + 1
}
