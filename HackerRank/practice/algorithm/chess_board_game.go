package main

import "fmt"

func main() {
	fmt.Println(chessboardGame(5, 2)) // expected: Second
	fmt.Println(chessboardGame(8, 8)) // expected: First
	fmt.Println(chessboardGame(5, 3)) // expected: First
}

func chessboardGame(x int32, y int32) string {
	x = x % 4
	y = y % 4
	if x == 3 || y == 3 || x == 0 || y == 0 {
		return "First"
	}
	return "Second"
}
