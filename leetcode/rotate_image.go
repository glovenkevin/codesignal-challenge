package main

// source: https://leetcode.com/problems/rotate-image/
func main() {
}

func rotate(matrix [][]int) {
	length := len(matrix) - 1
	if length == 0 {
		return
	}

	for i := 0; i <= length/2; i++ {
		matrix[i], matrix[length-i] = matrix[length-i], matrix[i]
	}

	for i := 0; i <= length; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
