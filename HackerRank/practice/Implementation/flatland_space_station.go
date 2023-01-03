package main

import (
	"fmt"
	"log"
	"sort"
)

// source: https://www.hackerrank.com/challenges/flatland-space-stations/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	pp := []struct {
		n int32
		c []int32
		a int32
	}{
		{n: 3, c: []int32{1}, a: 1},
		{n: 5, c: []int32{4, 0}, a: 2},
		{n: 6, c: []int32{0, 1, 2, 3, 4, 5}, a: 0},
		{n: 5, c: []int32{0}, a: 4},
		{n: 5, c: []int32{3}, a: 3},
		{n: 5, c: []int32{2}, a: 2},
		{n: 5, c: []int32{1, 3}, a: 1},
		{n: 5, c: []int32{1, 2}, a: 2},
		{n: 5, c: []int32{1, 2, 3}, a: 1},
	}
	for i, p := range pp {
		fmt.Printf("%d. Result: %d | Actual: %d \n", i+1, flatlandSpaceStations(p.n, p.c), p.a)
	}
}

// Solution: https://www.youtube.com/watch?v=6d1ZBC9Dj1Q
func flatlandSpaceStations(n int32, c []int32) int32 {
	var cc []int
	for _, node := range c {
		cc = append(cc, int(node))
	}
	sort.Ints(cc)

	var longest int32
	for i := 1; i < len(cc); i++ {
		longPath := (cc[i] - cc[i-1]) / 2
		if int32(longPath) > longest {
			longest = int32(longPath)
		}
	}

	if cc[0] > int(longest) {
		longest = int32(cc[0])
	}

	last := n - 1 - int32(cc[len(cc)-1])
	if last > longest {
		longest = last
	}

	return longest
}
