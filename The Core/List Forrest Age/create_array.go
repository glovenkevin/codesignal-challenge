package main

import "fmt"

func main() {
	fmt.Println(createArray(4))
}

func createArray(size int) []int {
	arr := []int{}
	for i := 0; i < size; i++ {
		arr = append(arr, 1)
	}
	return arr
}
