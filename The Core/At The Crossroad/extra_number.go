package main

import "fmt"

func main() {
	fmt.Println(extraNumber(2, 7, 2)) // expected: 7
}

func extraNumber(a int, b int, c int) int {
    if a == b {
        return c
    }
    if a == c {
        return b
    }
    return a
}

func extraNumber2(a int, b int, c int) int {
    return a ^ b ^ c
}