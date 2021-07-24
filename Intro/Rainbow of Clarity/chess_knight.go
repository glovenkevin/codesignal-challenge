package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(chessKnight("a1")) // expected: 2
	fmt.Println(chessKnight("c2")) // expected: 6
}

func chessKnight(cell string) int {
	rtn := 0
	rtn += findVertical(cell)
	rtn += findHorizontal(cell)
	return rtn
}

func findVertical(cell string) int {
	count := 0
	arrCell := strings.Split(cell, "")

	border := []rune("ah" + arrCell[0])
	xAxis := border[2]
	yAxis, _ := strconv.Atoi(arrCell[1])

	if yAxis+2 <= 8 &&
		yAxis-2 >= 1 {
		count += 2
		if xAxis-1 >= border[0] && // A
			xAxis+1 <= border[1] { // Z
			count += 2
		}
	} else if yAxis+2 <= 8 ||
		yAxis-2 >= 1 {
		count++
		if xAxis-1 >= border[0] && // A
			xAxis+1 <= border[1] { // Z
			count++
		}
	}
	return count
}

func findHorizontal(cell string) int {
	count := 0
	arrCell := strings.Split(cell, "")

	border := []rune("ah" + arrCell[0])
	xAxis := border[2]
	yAxis, _ := strconv.Atoi(arrCell[1])

	if xAxis-2 >= border[0] &&
		xAxis+2 <= border[1] {
		count += 2
		if yAxis-1 >= 1 && // A
			yAxis+1 <= 8 { // Z
			count += 2
		}
	} else if xAxis-2 >= border[0] ||
		xAxis+2 <= border[1] {
		count++
		if yAxis-1 >= 1 && // A
			yAxis+1 <= 8 { // Z
			count++
		}
	}
	return count
}
