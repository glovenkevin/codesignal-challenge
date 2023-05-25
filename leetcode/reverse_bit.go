package main

import (
	"fmt"
	"math/bits"
)

// source: https://leetcode.com/problems/reverse-bits/
func main() {
	tc := []struct {
		num uint32
		res uint32
	}{
		{
			num: 43261596,  // 00000010100101000001111010011100
			res: 964176192, // 00111001011110000010100101000000
		},
	}

	for _, t := range tc {
		res := reverseBits(t.num)

		if res != t.res {
			fmt.Println("test failed")
			continue
		}

		fmt.Println("test success")
	}
}

func reverseBits(num uint32) uint32 {
	var res uint32
	n := 31
	for n != -1 {
		last := num & 1
		last <<= n
		res += last
		num >>= 1
		n--
	}

	return res
}

func reverseBits2(num uint32) uint32 {
	return bits.Reverse32(num)
}
