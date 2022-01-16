package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(firstNotRepeatingCharacter("abacabad")) // expected: c
}

func firstNotRepeatingCharacter(s string) string {
	for {
		s = calcStr(s, s[:1])
		if len(s) == 1 {
			break
		}

		if s == "" {
			s = "_"
			break
		}
	}
	return s
}

func calcStr(str string, char string) string {
	regex := regexp.MustCompile(char)
	var size = len(regex.FindAllString(str, -1))

	if size > 1 {
		return regex.ReplaceAllString(str, "")
	}
	return char
}
