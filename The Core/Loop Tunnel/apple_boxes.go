package main

import "fmt"

func main() {
	fmt.Println(appleBoxes(5)) // result: -15
}

func appleBoxes(k int) int {
	rtn := 0
	flag := -1
	for i := 1; i <= k; i++ {
		rtn += (i * i) * flag
		flag *= -1
	}
	return rtn
}
