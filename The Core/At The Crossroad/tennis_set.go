package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(tennisSet(6, 3)) // expected: true
}

func tennisSet(score1 int, score2 int) bool {
	largest := math.Max(float64(score1), float64(score2))
	smallest := math.Min(float64(score1), float64(score2))

	return (largest == 6 && smallest < 5) || (largest == 7 && smallest >= 5 && smallest < 7)
}
