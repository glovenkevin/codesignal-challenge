package main

import "fmt"

func main() {
	fmt.Println(metroCard(31))
}

func metroCard(lastNumberOfDays int) []int {
	if lastNumberOfDays == 30 {
		return []int{31}
	} else if lastNumberOfDays == 31 {
		return []int{28, 30, 31}
	} else {
		return []int{31}
	}
}
