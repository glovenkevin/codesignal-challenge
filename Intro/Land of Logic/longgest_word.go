package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(longestWord("Ready, steady, go!"))
}

func longestWord(text string) string {
	pattern := "[^a-zA-Z]+"
	re := regexp.MustCompile(pattern)
	newStr := re.ReplaceAllString(text, " ")
	arrWord := strings.Split(newStr, " ")
	longgestWord := ""
	for _, val := range arrWord {
		if len(val) > len(longgestWord) {
			longgestWord = val
		}
	}

	return longgestWord
}
