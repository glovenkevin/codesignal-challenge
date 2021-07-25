package main

import "fmt"

func main() {
	fmt.Println(phoneCall(3, 1, 2, 20)) // expected: 14
}

func phoneCall(min1 int, min2_10 int, min11 int, s int) int {
	count := 0
	for s != 0 {
		if count == 0 {
			s -= min1
		} else if count < 10 {
			s -= min2_10
		} else {
			s -= min11
		}
		if s < 0 {
			break
		}
		count++
	}
	return count
}
