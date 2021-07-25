package main

import "fmt"

func main() {
	fmt.Println(circleOfNumbers(10, 2))  // expected: 7
	fmt.Println(circleOfNumbers2(10, 2)) // expected: 7
}

func circleOfNumbers(n int, firstNumber int) int {
	pos := (n / 2) + firstNumber
	if pos > n {
		pos -= n
	} else if pos == n {
		pos = 0
	}
	return pos

}

func circleOfNumbers2(n int, firstNumber int) int {
	return ((n / 2) + firstNumber) % n

}
