package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the next prize")) // Expected: panagram
	fmt.Println(pangrams("We promptly judged antique ivory buckles for the prize"))      // Expected: not panagram
}

func pangrams(s string) string {
	s = strings.ToLower(s)

	words := make(map[rune]bool)
	for _, rune := range s {
		if words[rune] == false {
			words[rune] = true
		}
	}

	if len(words) != 27 {
		return "not pangram"
	}
	return "pangram"
}
