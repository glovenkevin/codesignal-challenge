package main

import (
	"fmt"
	"math"
)

func main() {
	param1 := "ABCDEF"
	param2 := "FBDAMN"
	fmt.Println(commonChild(param1, param2)) // expected: 2 -> BD

	param1 = "SHINCHAN"
	param2 = "NOHARAAA"
	fmt.Println(commonChild(param1, param2)) // expected: 3 -> NHA
}

/*
 * Complete the 'commonChild' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func commonChild(s1 string, s2 string) int32 {

	var (
		sizeS1 int     = len(s1)
		sizeS2 int     = len(s2)
		set    [][]int = make([][]int, sizeS1+1)
	)
	for i := range set {
		set[i] = make([]int, sizeS2+1)
	}

	for pos1, r1 := range s1 {
		for pos2, r2 := range s2 {

			if r1 == r2 {
				set[pos1+1][pos2+1] = set[pos1][pos2] + 1
			}

			if r1 != r2 {
				set[pos1+1][pos2+1] = int(math.Max(float64(set[pos1+1][pos2]), float64(set[pos1][pos2+1])))
			}

		}
	}

	return int32(set[len(s1)][len(s2)])
}
