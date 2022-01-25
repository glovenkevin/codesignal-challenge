package main

import "fmt"

func main() {
	// fmt.Println(romanToInt("III"))     // Expected: 3
	// fmt.Println(romanToInt("LVIII"))   // Expected: 58
	// fmt.Println(romanToInt("MCMXCIV")) // Expected: 1994
	fmt.Println(romanToInt("MDCXCV")) // Expected: 1690
}

func romanToInt(s string) int {

	mapSymbol := make(map[string]int)
	mapSymbol["I"] = 1
	mapSymbol["V"] = 5
	mapSymbol["X"] = 10
	mapSymbol["L"] = 50
	mapSymbol["C"] = 100
	mapSymbol["D"] = 500
	mapSymbol["M"] = 1000

	total := 0
	ss := []rune(s)
	size := len(ss)
	for i := 0; i < size; i++ {
		currSize := mapSymbol[string(ss[i])]

		if i+1 != size {
			nextSize := mapSymbol[string(ss[i+1])]

			if currSize < nextSize {
				total += (nextSize - currSize)
				i++
			} else {
				total += currSize
			}

		} else {
			total += currSize
		}
	}

	return total
}
