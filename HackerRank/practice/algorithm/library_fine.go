package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(libraryFine(9, 6, 2015, 6, 6, 2015)) // Expected: 45
}

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	r := parseDate(d1, m1, y1)
	d := parseDate(d2, m2, y2)

	if r.Before(d) || r.Equal(d) {
		return 0
	}

	if r.Month() == d.Month() && r.Year() == d.Year() {
		return int32(r.Day()-d.Day()) * 15
	}

	if r.Year() == d.Year() {
		return int32(r.Month()-d.Month()) * 500
	}

	return int32(r.Year()-d.Year()) * 10000
}

func parseDate(d int32, m int32, y int32) time.Time {
	t := time.Date(int(y), time.Month(m), int(d), 0, 0, 0, 0, time.UTC)
	return t
}
