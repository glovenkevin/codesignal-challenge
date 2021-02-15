package main

import (
	"fmt"
)

func main() {
	fmt.Println(depositProfit(100, 20, 170)) // 3
	fmt.Println(depositProfit2(100, 20, 170)) // 3
}

func depositProfit(deposit int, rate int, threshold int) int { // Exceed time Limit
	counter := 0
	amount := float32(deposit)
	for {
        counter++
		amount += amount * ( float32(rate) / 100.0)
		if (amount >= float32(threshold)) {
            return counter
		}
	}
}

func depositProfit2(deposit int, rate int, threshold int) int {
	counter := 0
	depo := float32(deposit)
	for {
		counter++
        depo *= (float32(rate) + 100.0) / 100.0
		if (depo > float32(threshold)) {
			return counter
		}
	}
}