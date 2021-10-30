package main

import "fmt"

func main() {
	param := []int{1, 4, 6, 12}
	fmt.Println(firstReverseTry(param))
}

func firstReverseTry(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}

	arr[0], arr[len(arr)-1] = arr[len(arr)-1], arr[0]
	return arr
}
