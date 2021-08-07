package main

import "fmt"

func main() {
	fmt.Println(willYou(true, true, true)) // expected: false
}

func willYou(young bool, beautiful bool, loved bool) bool {
	return (young && beautiful && !loved) || (loved && (!young || !beautiful))
}
