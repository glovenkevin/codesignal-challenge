package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(validTime("13:58")) // expected: true
	fmt.Println(validTime("25:51")) // expected: false
}

func validTime(time string) bool {
	pattern := "([0-1][0-9]|[2][0-3]):([0-5][0-9])"
	regex := regexp.MustCompile(pattern)
	res := regex.MatchString(time)
	return res
}
