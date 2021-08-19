package main

import "fmt"

func main() {
	fmt.Println(magicalWell(1, 2 , 2)) // expected: 8
}

func magicalWell(a int, b int, n int) int {
    count := 0
    for n > 0 {
        count += a * b
        a++
        b++
        n--
    }
    return count
}
