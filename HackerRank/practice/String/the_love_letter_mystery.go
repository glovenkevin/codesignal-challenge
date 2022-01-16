package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(theLoveLetterMystery("abcd"))  // Expected: 4
	fmt.Println(theLoveLetterMystery("cba"))   // Expected: 2
	fmt.Println(theLoveLetterMystery("abc"))   // Expected: 2
	fmt.Println(theLoveLetterMystery("abcba")) // Expected: 0
}

/*
	To do this, he follows two rules:
	He can only reduce the value of a letter by , i.e. he can change d to c, but he cannot change c to d or d to b.
	The letter  may not be reduced any further.
*/
func theLoveLetterMystery(s string) int32 {

	var (
		totalStep int32  = 0
		ss        []rune = []rune(s)
	)

	for i, j := 0, len(ss)-1; i <= j; i, j = i+1, j-1 {

		if ss[i] != ss[j] {

			totalStep += int32(math.Abs(float64(ss[i]) - float64(ss[j])))

		}

	}

	return totalStep

}
