package main

import (
	"fmt"
)

// source: https://www.hackerrank.com/challenges/save-the-prisoner/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
func main() {
	pp := []struct {
		n   int32
		m   int32
		s   int32
		res int32
	}{
		{n: 5, m: 2, s: 1, res: 2},
		{n: 4, m: 6, s: 2, res: 3},
		{n: 4, m: 12, s: 1, res: 4},
	}

	for i, p := range pp {
		fmt.Printf("%d. Result: %d -> Actual: %d \n", i+1, saveThePrisoner(p.n, p.m, p.s), p.res)
	}
}

// n -> prisoner
// m -> total candy
// s -> start at index
func saveThePrisoner(n int32, m int32, s int32) int32 {
	sisaPermen := m % n
	index := s + sisaPermen - 1
	if index > n {
		index -= n
	}
	if index == 0 {
		index = n
	}
	return index
}
