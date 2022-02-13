package main

import (
	"fmt"
	"math/bits"
)

func main() {
	// fmt.Println(maximizingXor(10, 15))  // Expected: 7

	fmt.Println(maximizingXor2(10, 15)) // Expected: 7
	// fmt.Println(maximizingXor2(11, 100)) // Expected: 127
}

func maximizingXor(l int32, r int32) int32 {

	if l == r {
		return l ^ r
	}

	var larger int32 = 0
	for i := l; i <= r-1; i++ {
		for j := i + 1; j <= r; j++ {

			xorResult := i ^ j
			if xorResult >= larger {
				larger = xorResult
			}

		}
	}

	return larger
}

func maximizingXor2(l int32, r int32) int32 {

	xored := l ^ r
	fmt.Println("Xored: " + fmt.Sprintf("%08b", xored))
	fmt.Printf("Leading Zero: %d \n", bits.LeadingZeros32(uint32(xored)))

	sb := 32 - bits.LeadingZeros32(uint32(xored))

	fmt.Println("sb: " + fmt.Sprintf("%08b", (sb)))
	fmt.Println("1 << sb: " + fmt.Sprintf("%08b", 1<<(sb)))
	fmt.Println("result: " + fmt.Sprintf("%08b", (1<<(sb))-1))
	result := (1 << sb) - 1

	return int32(result)
}
