package main

import (
	"fmt"
)

func main() {
	// fmt.Println(sillyGame(1))    // Bob
	// fmt.Println(sillyGame(2))    // Allice
	// fmt.Println(sillyGame(5))    // Allice
	// fmt.Println(sillyGame(100))  // Allice
	// fmt.Println(sillyGame(996))  // Bob
	// fmt.Println(sillyGame(999))  // Bob
	// fmt.Println(sillyGame(1000)) // Allice

	fmt.Println(PrimeNumber(1000))  // 168
	fmt.Println(PrimeNumber(10000)) // 1229
}

// p1: Allice
// p2: Bob
// https://www.hackerrank.com/challenges/alice-and-bobs-silly-game/problem
func sillyGame(n int64) string {
	if n <= 1 {
		return "Bob"
	}

	p := PrimeNumber(n)
	fmt.Printf("n: %d, prime: %d, result: ", n, p)

	if p%2 == 0 {
		return "Bob"
	}
	return "Alice"
}

// Time space O(n)
// Space Complexity O(n*n)
// Method: Segmented sieve of erasthosthenes
func PrimeNumber(n int64) int {
	mp := map[int]bool{}
	t := 0
	for i := 2; i <= int(n); i++ { // O(n)
		_, ok := mp[i] // O(1)
		if !ok {
			t++
			mp[i] = true
			for j := i * 2; i*i <= int(n) && j <= int(n); j += i { // o(N)
				mp[j] = false
			}
		}
	}

	return t
}
