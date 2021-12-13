package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println(substrCount(5, "aaaa")) // Expected: 10
	// fmt.Println(substrCount(5, "asasd")) // Expected: 7
	// fmt.Println(substrCount(7, "abcbaba")) // Expected: 10
	// fmt.Println(substrCount(4, "aaaa")) // Expected: 10
	fmt.Println(substrCount(8, "mnonopoo")) // Expected: 12
}

/*
	RULE SPECIAL STRING:
		- All of the characters are the same, e.g. aaa.
		- All characters except the middle one are the same, e.g. aadaa.
*/

type point struct {
	key   rune
	count int
}

func substrCount(n int32, s string) int64 {

	var (
		totalCount  int     = 0
		count       int     = 0
		pp          []point = []point{}
		sizePp      int
		idx         int = 1
		currentRune rune
	)

	for _, r := range s {

		if currentRune == 0 {
			currentRune = r
			count = 1
			continue
		}

		if currentRune == r {
			count++
		}

		if currentRune != r {
			totalCount += (count * (count + 1) / 2)
			p := point{key: currentRune, count: count}
			pp = append(pp, p)

			currentRune = r
			count = 1
		}
	}

	totalCount += (count * (count + 1) / 2)
	p := point{key: currentRune, count: count}
	pp = append(pp, p)
	sizePp = len(pp)

	for idx < sizePp-1 {

		if pp[idx-1].key == pp[idx+1].key && pp[idx].count == 1 {
			totalCount += int(math.Min(float64(pp[idx-1].count), float64(pp[idx+1].count)))
		}

		idx++
	}

	return int64(totalCount)
}
