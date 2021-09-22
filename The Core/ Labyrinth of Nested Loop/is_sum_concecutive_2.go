package main

import "fmt"

func main() {
	fmt.Println(isSumOfConsecutive2(9))  // expected: 2
	fmt.Println(isSumOfConsecutive2(8))  // expected: 0
	fmt.Println(isSumOfConsecutive2(15)) // expected: 3
}

func isSumOfConsecutive2(n int) int {
	count := 0
	for i := 1; i < n; i++ {
		total := i
		counter := i + 1
		for total < n {
			total += counter
			if total == n {
				count++
				break
			}
			counter++
		}
	}
	return count
}
