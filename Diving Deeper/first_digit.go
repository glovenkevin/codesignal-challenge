package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(firstDigit("var_1__Int")) // Return '1'
}

func firstDigit(inputString string) string {
    var rtn string
    for _, val := range strings.Split(inputString, "") {
        if _, err := strconv.Atoi(val); err==nil {
            rtn = val
            break
        }
    }
    return rtn
}
