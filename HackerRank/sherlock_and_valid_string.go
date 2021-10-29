package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isValid("aabbcd"))            // expected: NO
	fmt.Println(isValid("abcdefghhgfedecba")) // expected: YES
	fmt.Println(isValid("aabbccddeefghi"))    // expected: NO

}

// string to be valid if all characters of the string appear the same number of times.
// It is also valid if he can remove just  character at  index in the string.
func isValid(s string) string {

	mapStr := make(map[rune]int)
	for _, rune := range s {
		mapStr[rune]++
	}

	mapFreq := make(map[int]int)
	for _, val := range mapStr {
		mapFreq[val]++
	}

	sizeFreq := len(mapFreq)
	if sizeFreq == 1 {
		return "YES"
	}

	if sizeFreq > 2 {
		return "NO"
	}

	arrKey := [2]int{}
	index := 0
	for key := range mapFreq {
		arrKey[index] = key
		index++
	}

	if math.Abs(float64(arrKey[0]-arrKey[1])) == 1 && (mapFreq[arrKey[0]] == 1 || mapFreq[arrKey[1]] == 1) {
		return "YES"
	}

	if (arrKey[0] == 1 && mapFreq[arrKey[0]] == 1) || arrKey[1] == 1 && mapFreq[arrKey[1]] == 1 {
		return "YES"
	}

	return "NO"
}
