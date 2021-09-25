package main

import "fmt"

func main() {
	fmt.Println(comfortableNumbers(10, 12)) // expected: 2
}

func comfortableNumbers(l int, r int) int {
	if l == r {
		return 0
	}

	rtn := 0
	for i := l; i < r; i++ {
		for j := l + 1; j <= r; j++ {
			if i == j {
				continue
			}
			if isComfort(i, j) && isComfort(j, i) {
				rtn++
			}
		}
	}
	return rtn
}

func isComfort(a int, b int) bool {
	return b <= a+sumDigit(a) && b >= a-sumDigit(a)
}

func sumDigit(val int) int {
	rtn := 0
	for val != 0 {
		rtn += val % 10
		val /= 10
	}
	// fmt.Println("Sum: ", rtn)
	return rtn
}
