package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(additionWithoutCarrying(456, 1734)) // Result: 1180
}

func additionWithoutCarrying(param1 int, param2 int) int {
	p1 := strconv.Itoa(param1)
	p2 := strconv.Itoa(param2)

	result := ""
	if len(p1) >= len(p2) {
		diff := len(p1) - len(p2)
		sufix := p1[diff:]
		result = p1[:diff] + addition(sufix, p2)
	} else {
		diff := len(p2) - len(p1)
		sufix := p2[diff:]
		result = p2[:diff] + addition(sufix, p1)
	}
	intResult, _ := strconv.Atoi(result)
	return intResult
}

func addition(param1 string, param2 string) string {
	rtn := ""
	size := len(param1)
	for i := 0; i < size; i++ {
		num1, _ := strconv.Atoi(param1[i : i+1])
		num2, _ := strconv.Atoi(param2[i : i+1])
		rtn += strconv.Itoa((num1 + num2) % 10)
	}
	return rtn
}
