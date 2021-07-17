package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(lineEncoding("aabbbc")) // expected: 2a3bc
}

func lineEncoding(s string) string {
    arrString := []rune(s)
    
    runeChar := arrString[0]
    tempCount := 0
    rtn := ""
    for _, val := range arrString {
        if runeChar == val {
            tempCount++
        } else {
            rtn += makeString(runeChar, tempCount)
            runeChar = val
            tempCount = 1
        }
    }
    rtn += makeString(runeChar, tempCount)
    
    return rtn
}

func makeString(s rune, count int) string {
    if count > 1 {
        return strconv.Itoa(count) + string(s)
    } else {
        return string(s)
    }
}