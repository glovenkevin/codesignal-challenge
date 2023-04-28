package main

import (
	"fmt"
	"math"
)

func main() {
	tc := []struct {
		x   int
		res int
	}{
		{4, 2},
		{8, 2},
		{3, 1},
		{9, 3},
	}

	for _, t := range tc {
		r := mySqrt(t.x)

		if r != t.res {
			fmt.Println("test fail")
			continue
		}

		fmt.Println("test successfull")
	}
}

// babylonian formula for finding square roots
func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}

	var (
		r = float64(x)
	)
	for i := 0; i <= 50; i++ {
		r = (r + (float64(x) / r)) / 2
	}

	return int(math.Floor(r))
}
