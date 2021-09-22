package main

import "fmt"

func main() {
	param := []int{6, 2, 3, 8}
	fmt.Println(makeArrayConsecutive2(param)) // expected: 3
}

func makeArrayConsecutive2(statues []int) int {
	statueLen := len(statues)
	biggest := 0
	lowest := 99
	for _, val := range statues {
		if val > biggest {
			biggest = val
		}
		if val < lowest {
			lowest = val
		}
	}

	if biggest-lowest == statueLen {
		return lowest
	}
	return (biggest - lowest + 1) - statueLen
}
