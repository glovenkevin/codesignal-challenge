package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(digitsProduct(450)) // expected: 2559
	fmt.Println(digitsProduct(12))  // expected: 26
}

func digitsProduct(product int) int {
	if product == 0 {
		return 10
	} else if product < 10 {
		return product
	}

	origin := product
	newSmallPos := 0
	newLength := 0
	for product != 1 {
		divider := 9
		for divider > 1 {
			if product%divider == 0 {
				nextDigit := product / divider
				newSmallPos = (divider * (int(math.Pow10(newLength)))) + newSmallPos
				newLength++
				product = nextDigit
				break
			}
			divider--
		}
		if product == origin {
			newSmallPos = -1
			break
		} else {
			origin = product
		}
	}

	return newSmallPos
}
