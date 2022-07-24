package main

import "fmt"

func main() {
	fmt.Println(towerBreakers(1, 7)) // expected: 1
	fmt.Println(towerBreakers(3, 7)) // expected: 1

	fmt.Println(towerBreakers(2, 2)) // expected: 2
	fmt.Println(towerBreakers(1, 4)) // expected: 1
}

// key Point: if there is 1 or an even number of tower than player 2 Win, else player 1
func towerBreakers(n int32, m int32) int32 {
	if m == 1 || n%2 == 0 {
		return 2
	}
	return 1
}
