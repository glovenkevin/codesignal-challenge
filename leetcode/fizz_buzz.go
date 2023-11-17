package main

import "strconv"

// source: https://leetcode.com/problems/fizz-buzz/
func main() {
}

func fizzBuzz(n int) []string {
	nn := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			nn = append(nn, "FizzBuzz")
		} else if i%3 == 0 {
			nn = append(nn, "Fizz")
		} else if i%5 == 0 {
			nn = append(nn, "Buzz")
		} else {
			nn = append(nn, strconv.Itoa(i))
		}
	}
	return nn
}
