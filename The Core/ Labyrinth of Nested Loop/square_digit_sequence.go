package main

import "fmt"

func main() {
	fmt.Println(squareDigitsSequence(16))  // expected: 9
	fmt.Println(squareDigitsSequence(103)) // expected: 4
}

func squareDigitsSequence(a0 int) int {
	val := a0
	memoizeVal := make(map[int]string)
	memoizeVal[val] = "x"
	for {
		val = calcSum(val)
		fmt.Println("Val :", val)

		if memoizeVal[val] == "x" {
			break
		}
		memoizeVal[val] = "x"
	}
	return len(memoizeVal) + 1
}

func calcSum(val int) int {
	xTen := 1
	rtn := 0
	tempVal := val
	for xTen <= val {
		modRes := tempVal % 10
		rtn += (modRes * modRes)
		tempVal /= 10
		if tempVal < 10 {
			tempVal += 10
		}
		xTen *= 10
	}
	return rtn
}
