package main

import "fmt"

func main() {
	p1 := []int{2, 3, 2, 3, 4, 5}
	fmt.Println(removeArrayPart(p1, 2, 4)) // expected: [2, 3, 5]
}

func removeArrayPart(inputArray []int, l int, r int) []int {
	return append(inputArray[:l], inputArray[r+1:]...)
}
