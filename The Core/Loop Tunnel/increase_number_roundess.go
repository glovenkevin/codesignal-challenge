package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(increaseNumberRoundness(902200100)) // result: true
	fmt.Println(increaseNumberRoundness(11000))     // result: false
	fmt.Println(increaseNumberRoundness(888))       // result: false

	fmt.Println(increaseNumberRoundness2(902200100)) // result: true
	fmt.Println(increaseNumberRoundness2(11000))     // result: false
	fmt.Println(increaseNumberRoundness2(888))       // result: false
}

/*
   use case:
   1. no zero at all
   2. there is zero but not at the bottom of the line
   3. there is no zero can be swapped to increase number rounds
*/
func increaseNumberRoundness(n int) bool {
	nString := strconv.Itoa(n)
	length := len(nString) - 1

	containZero := strings.Contains(nString, "0")
	if !containZero {
		return false
	} else {
		groupOfZero := 0
		lastPosition := false
		isGroup := false
		for idx, val := range []rune(nString) {
			char := string(val)
			if char == "0" && !isGroup {
				groupOfZero++
				isGroup = true
			} else if char != "0" {
				isGroup = false
			}

			if char == "0" && idx == length {
				lastPosition = true
			}
		}

		if groupOfZero > 1 || lastPosition == false {
			return true
		} else {
			return false
		}
	}
}

func increaseNumberRoundness2(n int) bool {
	nString := strconv.Itoa(n)
	arrRune := []rune(nString)
	size := len(arrRune) - 1
	for i := 0; i < size; i++ {
		if string(arrRune[i]) == "0" && string(arrRune[i+1]) != "0" {
			return true
		}
	}
	return false
}
