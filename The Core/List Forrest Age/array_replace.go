package main

import "fmt"

func main() {
	param1 := []int{1, 2, 1}
	param2 := 1
	param3 := 5
	fmt.Println(arrayReplace(param1, param2, param3))
}

func arrayReplace(inputArray []int, elemToReplace int, substitutionElem int) []int {
	for idx, val := range inputArray {
		if val == elemToReplace {
			inputArray[idx] = substitutionElem
		}
	}
	return inputArray
}
