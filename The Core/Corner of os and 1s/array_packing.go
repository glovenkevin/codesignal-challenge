package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(arrayPacking([]int{24, 85, 0}))
}

func arrayPacking(a []int) int {
	strBits := ""
	for _, val := range a {
		bits := fmt.Sprintf("%08b", val)
		strBits = bits + strBits
	}
	val, _ := strconv.ParseInt(strBits, 2, 64)
	return int(val)
}
