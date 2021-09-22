package main

import "fmt"

func main() {
	fmt.Println(countBlackCells(3, 4)) // expected: 6
}

func countBlackCells(n int, m int) int {
	if n == 1 || m == 1 {
		return n * m
	} else if n == m {
		return n + (n-1)*2
	} else {
		return m + n + gcd(n, m) - 2
	}
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
