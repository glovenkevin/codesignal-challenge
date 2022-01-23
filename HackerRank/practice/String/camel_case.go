package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(camelcase("saveChangesInTheEditor"))  // expected: 5
	fmt.Println(camelcase2("saveChangesInTheEditor")) // expected: 5
}

func camelcase(s string) int32 {
	regex := regexp.MustCompile("[A-Z]{1}")
	matched := regex.FindAllString(s, -1)
	return int32(len(matched) + 1)
}

func camelcase2(s string) int32 {
	regex := regexp.MustCompile("[A-Z]")
	newS := regex.ReplaceAllString(s, "")
	return int32(len(s) - len(newS) + 1)
}
