package main

import "fmt"

// source: https://leetcode.com/problems/spiral-matrix-ii
func main() {
	tc := []struct {
		n   int
		res [][]int
	}{
		{
			n: 1,
			res: [][]int{
				{1},
			},
		},
		{
			n: 2,
			res: [][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			n: 3,
			res: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			n: 5,
			res: [][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
		},
	}
	for _, c := range tc {
		res := generateMatrix(c.n)

	checkLoop:
		for i := range res {
			for j := range res[i] {
				if res[i][j] != c.res[i][j] {
					fmt.Printf("Result: %v, expected: %v \n", res, c.res)
					break checkLoop
				}
			}
		}

		fmt.Println("Test successful")
	}
}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	// anchor: 1:1, 2:2, 3:3, ...
	var anchor int = 0
	/*
		cl: column go to the left
		cr: column go to the right
		ru: row go up
		rd: row go down
	*/
	var pos string = "cr"
	var row, col, counter, total int = 0, 0, 1, n * n
	for total > 0 {
		matrix[row][col] = counter
		counter++
		total--

		if row == anchor && col == n-1-anchor { // 0, 4
			pos = "rd"
		}
		if row == n-1-anchor && col == n-1-anchor { // 4,4
			pos = "cl"
		}
		if row == n-1-anchor && col == anchor { // 4,0
			pos = "ru"
		}

		if pos == "cl" {
			col--
			continue
		}
		if pos == "cr" {
			col++
			continue
		}

		if pos == "rd" {
			row++
			continue
		}
		if pos == "ru" {
			row--
		}

		if row == anchor && col == anchor { // 0,0
			anchor++
			row = anchor
			col = anchor
			pos = "cr"
		}
	}

	return matrix
}
