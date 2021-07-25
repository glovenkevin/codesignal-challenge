package main

import "fmt"

func main() {
	fmt.Println(seatsInTheater(16, 11, 5, 3)) // expected: 96
}

func seatsInTheater(nCols int, nRows int, col int, row int) int {
	return (nCols - (col - 1)) * (nRows - row)
}
