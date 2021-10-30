package main

import "fmt"

func main() {
	fmt.Println(rounders(1234))  // result: 1000
	fmt.Println(rounders2(1234)) // result: 1000
}

func rounders(n int) int {
	x, y := 10, 1
	carry := 0
	rtn := 0
	for n/y != 0 {
		digit := ((n % x) / y) + carry
		if digit >= 5 {
			carry = 1
		} else {
			carry = 0
		}

		if x > n {
			rtn = digit * y
		}

		x *= 10
		y *= 10

	}
	return rtn
}

func rounders2(n int) int {
	x := 1
	for n > 10 {
		n = (n + 5) / 10
		x *= 10
	}
	return (n * x)
}
