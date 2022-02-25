package main

import "fmt"

func main() {
	fmt.Println(chocolateFeast(15, 3, 2)) // expected: 9
}

func chocolateFeast(n int32, c int32, m int32) int32 {
	var (
		choco int32 = n / c
	)

	return choco + chocoWrappers(choco, m, 0)
}

func chocoWrappers(w int32, m int32, sw int32) int32 {

	var (
		choco int32 = (w + sw) / m
		sisa  int32 = (w + sw) % m
	)

	if choco == 0 {
		return 0
	} else {
		return choco + chocoWrappers(choco, m, sisa)
	}
}
