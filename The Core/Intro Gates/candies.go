package main

import "fmt"

func main() {
	fmt.Println(candies(3, 10)) //expected: 9
}

func candies(n int, m int) int {
	return (m / n) * n
}
