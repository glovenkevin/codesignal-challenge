package main

import "fmt"

func main() {
	param := []int{2, 1, 3, 5, 3, 2}
	fmt.Println(firstDuplicate((param))) // expected: 3
}

func firstDuplicate(a []int) int {
	length := len(a)
	m := make(map[int]bool)
	n := make(map[int]int)
	for i := 0; i < length; i++ {
		if m[a[i]] == true {
			return a[i]
		} else {
			m[a[i]] = true
			n[a[i]] = i
		}
	}
	return -1
}
