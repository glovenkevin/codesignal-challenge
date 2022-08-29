package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := [][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}
	fmt.Println(diagonalDifference(p1)) // expected: 2

	p2 := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	fmt.Println(diagonalDifference(p2)) // expected: 15
}

// Given a square matrix, calculate the absolute difference between the sums of its diagonals
// constraint: square matrix, size -100 <= arr[i][j] <= 100
func diagonalDifference(arr [][]int32) int32 {
	var (
		ln         = len(arr)
		l, r int32 = 0, 0
	)
	for i := 0; i < ln; i++ {
		l += arr[i][i]
		r += arr[i][ln-1-i]
	}

	fmt.Println("left", l)
	fmt.Println("right", r)
	return int32(math.Abs(float64(l - r)))
}
