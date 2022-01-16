package main

import "fmt"

func main() {
	fmt.Println(countingValleys(8, "UDDDUDUU")) // Expected: 1
}

// Time Complexity O(n)
func countingValleys(steps int32, path string) int32 {
	var (
		/*
		   Flag for current position:
		   -1 -> D / Down
		   0 -> Sea Level
		   1 -> U / Up
		*/
		pos          int   = 0
		countValleys int32 = 0
	)

	for _, rune := range path {
		char := string(rune)

		if char == "D" {
			pos -= 1
			continue
		}

		if char == "U" {
			pos += 1

			if pos == 0 {
				countValleys++
			}

			continue
		}

	}

	return countValleys
}
