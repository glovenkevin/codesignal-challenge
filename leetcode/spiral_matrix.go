package main

import "fmt"

// source: https://leetcode.com/problems/spiral-matrix/
func main() {
	tc := []struct {
		matrix [][]int
		res    []int
	}{
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			res: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			res: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			matrix: [][]int{
				{1, 2, 3},
			},
			res: []int{1, 2, 3},
		},
		{
			matrix: [][]int{
				{1},
				{2},
				{3},
			},
			res: []int{1, 2, 3},
		},
	}

test_loop:
	for _, t := range tc {
		res := spiralOrder(t.matrix)

		for i, v := range res {
			if v != t.res[i] {
				fmt.Println("test failed")
				continue test_loop
			}
		}
		fmt.Println("test success")
	}
}

func spiralOrder(matrix [][]int) []int {
	var (
		length, width = len(matrix[0]), len(matrix)
		matrixSize    = length * width
		ret           = make([]int, 0, matrixSize)

		mode     int = 0
		boundary int
		dir      = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		col, row int
	)

	for i := 0; i < matrixSize; i++ {
		ret = append(ret, matrix[row][col])

		if row-1 == boundary && col == boundary && mode == 3 {
			boundary++
			row = boundary
			col = boundary
			mode = 0
			continue
		}

		var isTopRight, isBottomRight, isBottomLeft bool
		if col+1 == length-boundary && row == boundary && mode == 0 {
			isTopRight = true
		}
		if row+1 == width-boundary && col+1 == length-boundary && mode == 1 {
			isBottomRight = true
		}
		if col == boundary && row+1 == width-boundary && mode == 2 {
			isBottomLeft = true
		}

		if isTopRight || isBottomRight || isBottomLeft {
			mode++
		}

		row += dir[mode][0]
		col += dir[mode][1]
	}

	return ret
}
