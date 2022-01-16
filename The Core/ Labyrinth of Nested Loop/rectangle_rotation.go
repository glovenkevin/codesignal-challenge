package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(rectangleRotation2(6, 4))   // expected: 23
	fmt.Println(rectangleRotation2(2, 2))   // expected: 5
	fmt.Println(rectangleRotation2(50, 50)) // expected: 2521
	fmt.Println(rectangleRotation2(30, 2))  // expected: 65
}

// Source answer: https://medium.com/hard-mode/coding-challenge-rectangle-rotation-10e2a2416ef3
func rectangleRotation(a int, b int) int {
	aHalfBisect := (float64(a) / math.Sqrt2) / 2
	bHalfBisect := (float64(b) / math.Sqrt2) / 2
	rect1 := [2]float64{math.Floor(aHalfBisect)*2 + 1, math.Floor(bHalfBisect)*2 + 1}
	total1 := int(rect1[0] * rect1[1])

	var rect2X float64
	if aHalfBisect-math.Floor(aHalfBisect) < 0.5 {
		rect2X = rect1[0] - 1
	} else {
		rect2X = rect1[0] + 1
	}

	var rect2Y float64
	if bHalfBisect-math.Floor(bHalfBisect) < 0.5 {
		rect2Y = rect1[1] - 1
	} else {
		rect2Y = rect1[1] + 1
	}
	total2 := int(rect2X * rect2Y)

	return total1 + total2
}

func rectangleRotation2(a int, b int) int {
	a = (int)(float64(a) / math.Sqrt2)
	b = (int)(float64(b) / math.Sqrt2)
	c := (a+1)*(b+1) + a*b

	if c%2 == 0 {
		return (c - 1)
	} else {
		return c
	}
}
