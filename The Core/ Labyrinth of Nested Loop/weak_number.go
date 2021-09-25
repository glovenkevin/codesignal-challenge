package main

import "fmt"

func main() {
	// fmt.Println(calcTotalDivisor(4)) // expected: 3
	// fmt.Println(calcTotalDivisor(9)) // expected: 3
	fmt.Println(weakNumbers(9)) // expected: [2, 2]
}

func weakNumbers(n int) []int {
	rtn := [2]int{0, 0}
	totalDiv := make(map[int]int)
	for i := 1; i <= n; i++ {
		divisor := calcTotalDivisor(i)
		totalDiv[i] = divisor
		weaknesLevel := 0
		for j := 1; j < i; j++ {
			if totalDiv[j] > divisor {
				weaknesLevel++
			}
		}

		if rtn[0] < weaknesLevel {
			rtn[0] = weaknesLevel
			rtn[1] = 1
		} else if rtn[0] == weaknesLevel {
			rtn[1]++
		}
	}
	return rtn[:]
}

func calcTotalDivisor(n int) int {
	rtn := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			rtn++
			if n/i != i {
				rtn++
			}
		}
	}
	return rtn
}
