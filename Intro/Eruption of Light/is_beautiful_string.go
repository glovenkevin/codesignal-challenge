package main

import "fmt"

func main() {
	fmt.Println(isBeautifulString("zaa"))
}

/*
	Rule:
	1. alphabet complete a-z
	2. alphabet is ordinal a-z
	3. current char can't have more than previous char
*/
func isBeautifulString(inputString string) bool {
	var rtn = true
	var arrayAlphabet = []rune("abcdefghijklmnopqrstuvwxyz")

	var mapString = countStringAlphabet(inputString)
	var countBefore = 99
	var countCurrent = 0
	var counter = 0
	for counter < len(arrayAlphabet) {
		countCurrent = mapString[int(arrayAlphabet[counter])]
		if counter == 0 && countCurrent == 0 {
			rtn = false
		} else if countBefore < countCurrent {
			rtn = false
		} else {
			countBefore = countCurrent
		}
		counter++
	}

	return rtn
}

func countStringAlphabet(input string) map[int]int {
	var rtn = make(map[int]int)
	var arrayInputString = []rune(input)
	var counter = 0
	for counter < len(arrayInputString) {
		var ascii = int(arrayInputString[counter])
		rtn[ascii] += 1
		counter++
	}

	return rtn
}
