package main

import "fmt"

func main() {
	fmt.Println(candles(5, 2)) // expected: 9
}

func candles(candlesNumber int, makeNew int) int {

	if candlesNumber < makeNew {
		return candlesNumber
	}

	rtn := 0 // 5
	leftOver := 0
	candle := candlesNumber // 5
	for candlesNumber > 0 { // 5 2 1 1
		rtn += candlesNumber
		candle = (candlesNumber + leftOver) / makeNew   // 2 1 1
		leftOver = (candlesNumber + leftOver) % makeNew // 1 1 0
		candlesNumber = candle
	}
	return rtn
}

func candels2(candlesNumber int, makeNew int) int {
	return candlesNumber + (candlesNumber-1)/(makeNew-1)
}
