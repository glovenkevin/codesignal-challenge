package main

import (
	"fmt"
	"strconv"
)

func main() {
	p1 := []int32{4, 5, 6, 7}
	fmt.Println(fairRationsForum(p1)) // Expected: 4

	p2 := []int32{2, 3, 4, 5, 6}
	fmt.Println(fairRationsForum(p2)) // Expected: 4

	p3 := []int32{1, 2}
	fmt.Println(fairRationsForum(p3)) // Expected: "NO"

	p4 := []int32{1, 1, 2, 1}
	fmt.Println(fairRationsForum(p4)) // Expected: "NO"

	p5 := []int32{1, 2, 1}
	fmt.Println(fairRationsForum(p5)) // Expected: 2
}

// Complexity: O(n)
func fairRations(B []int32) string {
	var odd int
	for _, b := range B {
		if b%2 == 1 {
			odd++
		}
	}
	if odd%2 != 0 {
		return "NO"
	}

	var i, b int32
	var l = int32(len(B))
	for i < l {
		if B[i]%2 == 0 {
			i++
			continue
		}

		B[i]++
		if i+1 == l {
			B[i-1]++
			i = 0
			continue
		}
		B[i+1]++
		b += 2
		i = 0
	}

	return strconv.Itoa(int(b))
}

// Forum Discussion Solution: O(n)
func fairRationsForum(B []int32) string {
	var res int32
	for i := range B[:len(B)-1] {
		if B[i]%2 == 1 {
			B[i]++
			B[i+1]++
			res += 2
		}
	}

	if B[len(B)-1]%2 == 1 {
		return "NO"
	}

	return strconv.Itoa(int(res))
}
