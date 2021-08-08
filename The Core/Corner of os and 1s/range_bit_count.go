package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(rangeBitCount(2, 7))
}

func rangeBitCount(a int, b int) int {
	rtn := 0
	for i := a; i <= b; i++ {
		bit := strconv.FormatInt(int64(i), 2)
		bit = strings.ReplaceAll(bit, "0", "")
		rtn += len(bit)
	}
	return rtn
}
