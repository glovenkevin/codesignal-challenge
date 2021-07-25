package main

import "fmt"

func main() {
	fmt.Println(lateRide(240)) // expected: 4
}

func lateRide(n int) int {
	hour := n / 60
	min := n % 60
	return (hour / 10) + (hour % 10) + (min / 10) + (min % 10)
}
