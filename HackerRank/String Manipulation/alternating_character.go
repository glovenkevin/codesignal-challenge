package main

import "fmt"

func main() {
	fmt.Println(alternatingCharacters("AAA"))    // expected: 2
	fmt.Println(alternatingCharacters("BBBB"))   // expected: 3
	fmt.Println(alternatingCharacters("ABAB"))   // expected: 0
	fmt.Println(alternatingCharacters("AAABBB")) // expected: 4
}

func alternatingCharacters(s string) int32 {

	var (
		tmpRune rune  = 0
		count   int32 = 0
	)
	for _, rune := range s {
		if tmpRune != rune {
			tmpRune = rune
		} else {
			count++
		}
	}
	return count
}
