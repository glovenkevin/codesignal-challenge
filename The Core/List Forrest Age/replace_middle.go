package main

import "fmt"

func main() {
	param := []int{7, 2, 2, 5, 10, 7}
	fmt.Println(replaceMiddle(param)) // expected: [7, 2, 7, 10, 7]
}

func replaceMiddle(arr []int) []int {
	arrLen := len(arr)
	if arrLen%2 == 0 {
		idx := arrLen / 2
		mid := arr[idx] + arr[idx-1]
		return append(append(arr[:idx-1], mid), arr[idx+1:]...)
	} else {
		return arr
	}
}
