package main

import "fmt"

func main() {
	a := []int{2, 4, 7}
	fmt.Println(absoluteValuesSumMinimization(a))

	b := []int{1, 1, 3, 4}
	fmt.Println(absoluteValuesSumMinimization(b))
}

func absoluteValuesSumMinimization(a []int) int {
	var middle int
	if (len(a) % 2 == 1) {
		return a[(len(a) - 1) / 2 ]
	} else {
		middle = len(a) / 2
		if (a[middle] > a[middle-1]) {
			return a[middle-1]
		} else {
			return a[middle]
		}
	}
}

func absoluteValuesSumMinimization2(a []int) int {
	return a[(len(a) - 1) / 2 ]
}