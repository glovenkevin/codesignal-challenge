package main

import "fmt"

func main() {
	fmt.Println(circleOfNumbers(10, 2)) // 7
	fmt.Println(circleOfNumbers(10, 7)) // 2

	fmt.Println(circleOfNumbers2(10, 2))
	fmt.Println(circleOfNumbers2(10, 7))
}

func circleOfNumbers(n int, firstNumber int) int {
	halfSize := (n / 2)
	if firstNumber >= halfSize {
		return firstNumber - halfSize
	} else {
		return firstNumber + halfSize
	}
}

func circleOfNumbers2(n int, firstNumber int) int {
	return (firstNumber + (n / 2)) % n
}