package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sumUpNumbers("2 apples, 12 oranges")) // expected: 14
	fmt.Println(sumUpNumbers("no digits at all"))     // expected: 0
}

func sumUpNumbers(inputString string) int {
	pattern := `[^\d]`
	regex := regexp.MustCompile(pattern)
	newStr := regex.ReplaceAllString(inputString, " ")
	arrString := strings.Split(newStr, " ")
	totalInt := 0
	for _, val := range arrString {
		if num, err := strconv.Atoi(val); err == nil {
			totalInt += num
		}
	}

	return totalInt
}
