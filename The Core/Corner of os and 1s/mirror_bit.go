package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(mirrorBits(97)) // expected: 67
}

func mirrorBits(a int) int {
	bitAngka := strconv.FormatInt(int64(a), 2)
	reverseBit := ""
	for _, val := range bitAngka {
		reverseBit = string(val) + reverseBit
	}
	rtn, _ := strconv.ParseInt(reverseBit, 2, 64)
	return int(rtn)
}
