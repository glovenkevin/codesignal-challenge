package main

import (
	"fmt"
)

func main() {
	fmt.print(buildPalindrome("abcdc")) // abcdcba
}

/* 
    Plan:
        1. check if pallindrom
        2. if not slice it become 2 and add it on behind
        3. back to number 1
*/
func buildPalindrome(st string) string {
    if (checkIfPallindrom(st)) {
        return st
    } else {
        var size = len(st) -1
        var counter = 0
        var newStr = ""
        for counter < size {
            newStr = st + getStringBefore(st, counter)
            if (checkIfPallindrom(newStr)) {
                break
            } else {
                counter++
            }
        }
        return newStr
    }
}

func checkIfPallindrom(val string) bool {
    var arrString = []rune(val)
    var reversedString = ""
    for i := len(arrString)-1; i >= 0; i-- {
        reversedString += string(arrString[i])
    }
    return val == reversedString
}

func getStringBefore(val string, to int) string {
    var arrString = []rune(val)
    var newStr = ""
    
    for i := to; i >= 0; i-- {
        newStr += string(arrString[i])
    }
    return newStr
}