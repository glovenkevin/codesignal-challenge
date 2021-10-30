package main

import "fmt"

func main() {
	param := []int{7, 2, 2, 5, 10, 7}
	fmt.Println(param) // expected: true
}

func isSmooth(arr []int) bool {
	arrLen := len(arr)
	mid := 0
	if arrLen%2 == 0 {
		idx := arrLen / 2
		mid = arr[idx] + arr[idx-1]
	} else {
		mid = arr[(arrLen-1)/2]
	}
	fmt.Println("mid: ", mid)
	return arr[0] == arr[arrLen-1] && arr[0] == mid
}
