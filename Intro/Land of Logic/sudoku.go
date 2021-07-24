package main

import (
	"fmt"
)

func main() {
	param := [][]int{
		{1, 3, 2, 5, 4, 6, 9, 8, 7},
		{4, 6, 5, 8, 7, 9, 3, 2, 1},
		{7, 9, 8, 2, 1, 3, 6, 5, 4},
		{9, 2, 1, 4, 3, 5, 8, 7, 6},
		{3, 5, 4, 7, 6, 8, 2, 1, 9},
		{6, 8, 7, 1, 9, 2, 5, 4, 3},
		{5, 7, 6, 9, 8, 1, 4, 3, 2},
		{2, 4, 3, 6, 5, 7, 1, 9, 8},
		{8, 1, 9, 3, 2, 4, 7, 6, 5},
	}
	fmt.Println(sudoku(param)) // Expected true

	param2 := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	fmt.Println(sudoku(param2)) // Excpected false

	// s+=grid[y-y%3+x/3][y%3*3+x%3]
	// y := 6
	// x := 2
	// fmt.Println(y - y%3 + x/3)
	// fmt.Println(y%3*3 + x%3)
	// fmt.Println(param[y-y%3+x/3][y%3*3+x%3])
}

func sudoku(grid [][]int) bool {
	loopX := 0
	totalY := [9]int{}
	for loopX < 9 {
		totalX := [3]int{}
		loopY := 0
		for loopY < 9 {

			totalNum := 0
			counter := 0
			flagNum := make(map[int]bool)
			for x := loopX; x < loopX+3; x++ {
				for y := loopY; y < loopY+3; y++ {
					if !flagNum[grid[x][y]] {
						flagNum[grid[x][y]] = true
						totalNum += grid[x][y]
						totalX[counter] += grid[x][y]
						totalY[y] += grid[x][y]
					} else {
						return false
					}

				}
				counter++
			}

			if totalNum != 45 {
				return false
			}

			loopY += 3
		}

		for _, val := range totalX {
			if val != 45 {
				return false
			}
		}

		loopX += 3
	}

	for _, val := range totalY {
		if val != 45 {
			return false
		}
	}

	return true
}
