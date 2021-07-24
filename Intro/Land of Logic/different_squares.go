package main

import "fmt"

func main() {
	param := [][]int{
		{1, 2, 1},
		{2, 2, 2},
		{2, 2, 2},
		{1, 2, 3},
		{2, 2, 1},
	}
	fmt.Println(differentSquares(param))
}

type rectangle struct {
	topLeft     int
	topRight    int
	bottomLeft  int
	bottomRight int
}

func differentSquares(matrix [][]int) int {
	arrRectagle := make(map[rectangle]bool)
	xLength := len(matrix)
	yLength := len(matrix[0])

	x := 0
	y := 0

	for x < xLength-1 {
		for y < yLength-1 {
			rect := rectangle{
				topLeft:     matrix[x][y],
				topRight:    matrix[x][y+1],
				bottomLeft:  matrix[x+1][y],
				bottomRight: matrix[x+1][y+1],
			}
			if !arrRectagle[rect] {
				arrRectagle[rect] = true
			}
			y++
		}
		x++
		y = 0
	}

	return len(arrRectagle)
}
