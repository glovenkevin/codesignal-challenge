package main

// source: https://leetcode.com/problems/number-of-1-bits
func main() {
}

func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		bit := num & 1
		if bit == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

func hammingWeightSimpler(num uint32) int {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		res += num & 1
		num = num >> 1
	}
	return int(res)
}
