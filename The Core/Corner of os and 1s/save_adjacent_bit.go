package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(swapAdjacentBits(13)) // expected: 14
	fmt.Println(swapAdjacentBits(74)) // expected: 133
}

func swapAdjacentBits(n int) int {
	return swap(n)
}

func swap(n int) int {
	bit := strconv.FormatInt(int64(n), 2)
	bit = fmt.Sprintf("%032s", bit)
	swappedBit := swapBit(bit)
	val, _ := strconv.ParseInt(swappedBit, 2, 64)
	return int(val)
}

func swapBit(n string) string {
	length := len(n) - 1
	arrStr := []rune(n)
	for idx := 0; idx < length; idx += 2 {
		arrStr[idx], arrStr[idx+1] = arrStr[idx+1], arrStr[idx]
	}
	return string(arrStr)
}
