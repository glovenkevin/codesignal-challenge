package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(findEmailDomain("\"very.unusual.@.unusual.com\"@usual.com"))
}

func findEmailDomain(address string) string {
	splitResult := strings.Split(address, "@")
	return splitResult[len(splitResult)-1]
}
