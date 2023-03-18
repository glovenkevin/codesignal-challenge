package main

import (
	"bytes"
	"fmt"
)

// source: https://leetcode.com/problems/integer-to-roman/
func main() {
	tc := []struct {
		num int
		res string
	}{
		{7, "VII"},
		{3, "III"},
		{2017, "MMXVII"},
	}
	for _, c := range tc {
		res := intToRoman(c.num)
		if res != c.res {
			fmt.Printf("res: %s | expected: %s \n", res, c.res)
			continue
		}
		fmt.Println("test successful")
	}
}

func intToRoman(num int) string {
	var (
		sortedIntRoman = [13]int{
			1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
		}
		romanMap = map[int]string{
			1000: "M",
			900:  "CM",
			500:  "D",
			400:  "CD",
			100:  "C",
			90:   "XC",
			50:   "L",
			40:   "XL",
			10:   "X",
			9:    "IX",
			5:    "V",
			4:    "IV",
			1:    "I",
		}
		bb bytes.Buffer
	)

	for _, v := range sortedIntRoman {
		if num < v {
			continue
		}

		for num >= v {
			num -= v
			bb.WriteString(romanMap[v])
		}
	}

	return bb.String()
}
