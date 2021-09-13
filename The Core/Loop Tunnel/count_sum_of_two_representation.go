package main

import "fmt"

func main() {
	fmt.Println(countSumOfTwoRepresentations2(24, 8, 16)) // expected: 4
	fmt.Println(countSumOfTwoRepresentations2(1000000, 490000, 900000))
	fmt.Println(countSumOfTwoRepresentations2(93, 24, 58))
}

func countSumOfTwoRepresentations2(n int, l int, r int) int {
	count := 0
	tempR := r
	for low := l; low <= tempR; low++ {
		if low+tempR > n && low < tempR {
			tempR--
		}
		if low+tempR == n && low <= tempR {
			count++
		}
	}

	if count == 0 {
		for high := r; high >= l; high-- {
			if l+high == n {
				count++
				l++
			}
		}
	}
	return count
}

// other simpler solve
func countSumOfTwoRepresentations23(n int, l int, r int) int {
	count := 0
	for i := l; i <= r; i++ {
		if n-i >= i && n-i <= r {
			count++
		}
	}
	return count
}
