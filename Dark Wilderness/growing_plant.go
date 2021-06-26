package main

import "fmt"

func main() {
	fmt.Println(growingPlant(100, 10, 910)) // Should be 10
}

func growingPlant(upSpeed int, downSpeed int, desiredHeight int) int {
    day := 1;
    height := upSpeed;
    for height < desiredHeight {
        height = (height-downSpeed) + upSpeed
        day++
    }
    return day
}