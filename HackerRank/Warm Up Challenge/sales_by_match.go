package main

import "fmt"

func main() {
	param := []int32{1, 1, 3, 1, 2, 1, 3, 3, 3, 3}
	fmt.Println(sockMerchant(10, param)) // Expected: 4
}

// time complex: O(n)
func sockMerchant(n int32, ar []int32) int32 {

	mapArr := make(map[int32]int32)
	var sum int32 = 0
	for _, a := range ar {
		mapArr[a] += 1

		if mapArr[a]/2 > 0 {
			sum += 1
			mapArr[a] -= 2
		}
	}

	return sum
}
