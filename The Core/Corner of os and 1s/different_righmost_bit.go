package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(differentRightmostBit(11, 13)) // expected: 2
	fmt.Println(differentRightmostBit(7, 23))  // expected: 16
}

func differentRightmostBit(a int, b int) int {
	c := a ^ b
	bit := strconv.FormatInt(int64(c), 2)
	arrStr := []rune(bit)
	idx := 0
	for i := len(bit) - 1; i >= 0; i-- {
		if string(arrStr[i]) == "1" {
			idx = len(bit) - (i + 1)
			break
		}
	}
	return pow2(idx)
}

func pow2(n int) int {
	rtn := 1
	for i := 0; i < n; i++ {
		rtn *= 2
	}
	return rtn
}
