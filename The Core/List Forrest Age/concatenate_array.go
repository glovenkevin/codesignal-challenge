package main

import "fmt"

func main() {
	param1 := []int{1, 2, 3}
	param2 := []int{4, 5, 6}
	fmt.Println(concatenateArrays(param1, param2))
}

func concatenateArrays(a []int, b []int) []int {
	return append(a, b...)
}
