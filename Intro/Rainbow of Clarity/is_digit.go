package main

import (
	"regexp"
	"fmt"
)

func main() {
	fmt.Println(isDigit("1")) // true
	fmt.Println(isDigit("-")) // false
}

func isDigit(symbol string) bool {
    var pattern = "[\\d]{1}"
    var regex = regexp.MustCompile(pattern)
    return nil != regex.FindStringSubmatch(symbol)
}
