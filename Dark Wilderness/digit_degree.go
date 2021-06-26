package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(digitDegree(10)) // Output 1
}

func digitDegree(n int) int {
    if n<10 {
        return 0
    } else {
        count := 0
        for n >= 10 {
            temp := 0
            for _, val := range strconv.Itoa(n) {
                i , _ := strconv.Atoi(string(val))
                temp += i
            }
            n = temp
            count++
        }
        return count
    }
}