package main

import "fmt"

func main() {

	aa := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	aa = [][]int{
		{40, 12, 15, 37, 33, 11, 45, 13, 25, 3},
		{37, 35, 15, 43, 23, 12, 22, 29, 46, 43},
		{44, 19, 15, 12, 30, 2, 45, 7, 47, 6},
		{48, 4, 40, 10, 16, 22, 18, 36, 27, 48},
		{45, 17, 36, 28, 47, 46, 8, 4, 17, 3},
		{14, 9, 33, 1, 6, 31, 7, 38, 25, 17},
		{31, 9, 17, 11, 29, 42, 38, 10, 48, 6},
		{12, 13, 42, 3, 47, 24, 28, 22, 3, 47},
		{38, 23, 26, 3, 23, 27, 14, 40, 15, 22},
		{8, 46, 20, 21, 35, 4, 36, 18, 32, 3},
	}

	fmt.Println(rotateImage(aa))
}

// Try to solve this task in-place (with O(1) additional memory)
// Guarantee n*n matrix
func rotateImage(aa [][]int) [][]int {
	length := len(aa) - 1
	if length == 0 {
		return aa
	}

	for i := 0; i <= length/2; i++ {
		aa[i], aa[length-i] = aa[length-i], aa[i]
	}
	fmt.Println(aa)

	for i := 0; i <= length; i++ {
		for j := 0; j <= i; j++ {
			aa[i][j], aa[j][i] = aa[j][i], aa[i][j]
		}
	}

	return aa
}
