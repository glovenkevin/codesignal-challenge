package main

import "fmt"

// source: https://leetcode.com/problems/pascals-triangle
func main() {
	fmt.Println(generate(1)) // [[1]]
	fmt.Println(generate(3)) // [[1], [1, 2, 1], [1, 3, 3 , 1]]
	fmt.Println(generate(4)) // [[1], [1, 2, 1], [1, 3, 3 , 1], [1, 4, 6, 4, 1]]
}

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	ret[0] = []int{1}
	if numRows == 1 {
		return ret
	}

	for i := 1; i < numRows; i++ {
		ret[i] = make([]int, i+1)
		for j := 0; j < len(ret[i]); j++ {
			if j == 0 || j == i {
				ret[i][j] = 1
				continue
			}

			ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
		}
	}

	return ret
}
