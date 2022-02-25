package main

import "fmt"

func main() {
	param := []int32{1, 4, 3, 2}
	fmt.Println(reverseArray(param)) // Expected: 2, 3, 4, 1
}

func reverseArray(a []int32) []int32 {
	al := len(a)

	for i, j := 0, al-1; i <= j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}

	return a
}
