package main

import "fmt"

func main() {
	fmt.Println(arithmeticExpression(6, 2, 3)) // expected: true
}

func arithmeticExpression(a int, b int, c int) bool {
	return a+b == c ||
		a-b == c ||
		a*b == c ||
		float32(a)/float32(b) == float32(c)
}

func arithmeticExpression2(a int, b int, c int) bool {
	return a+b == c ||
		a-b == c ||
		a*b == c ||
		a == b*c
}
