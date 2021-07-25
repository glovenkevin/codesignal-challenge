package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(addTwoDigits(29)) // expected: 11
}

func addTwoDigits(n int) int {
	arrStr := strings.Split(strconv.Itoa(n), "")
	obj1, _ := strconv.Atoi(arrStr[0])
	obj2, _ := strconv.Atoi(arrStr[1])
	return obj1 + obj2
}

func addTwoDigits2(n int) int {
	return (n / 2) + (n % 2)
}
