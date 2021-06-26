package main

import (
    "unicode"
	"fmt"
)

func main() {
	fmt.Println(longestDigitsPrefix("123aa")) // print 123
}

func longestDigitsPrefix(inputString string) string {
    rtn := ""
    for _ , char := range inputString {
        if !unicode.IsDigit(char) {
            return rtn
        } else {
            rtn += string(char)
        }
    }
    return rtn
}
