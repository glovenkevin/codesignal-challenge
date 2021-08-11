package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(secondRightmostZeroBit(37))
}

func secondRightmostZeroBit(n int) int {
	return calc(n)
}

func calc(n int) int {
	pos := getZero(n) - 1
	rtn := 2
	for pos > 0 {
		rtn *= 2
		pos--
	}
	return rtn
}

func getZero(n int) int {
	bit := strconv.FormatInt(int64(n), 2)
	first := false
	arrBit := []rune(bit)
	length := len(bit) - 1
	pos := 0
	for i := length; i >= 0; i-- {
		if string(arrBit[i]) == "0" && !first {
			first = true
		} else if string(arrBit[i]) == "0" {
			pos = i
			break
		}
	}
	return length - pos
}
