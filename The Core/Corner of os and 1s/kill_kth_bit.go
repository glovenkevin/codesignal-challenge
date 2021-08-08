package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(killKthBit(37, 3))
}

func killKthBit(n int, k int) int {
	bit := strconv.FormatInt(int64(n), 2)
	pos := len(bit) - k

	tempBit := ""
	for i, rune := range bit {
		if i == pos {
			tempBit += "0"
		} else {
			tempBit += string(rune)
		}
	}

	val, _ := strconv.ParseInt(tempBit, 2, 64)

	return int(val)
}
