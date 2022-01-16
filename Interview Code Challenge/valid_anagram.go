package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValidAnagram("danger", "garden")) // Expected: True

	rr := []rune("danger")
	insertionSort(rr, len(rr))
	fmt.Println(string(rr))

	rrInt := []int{5, 4, 2, 1, 3}
	insertionSortInt(rrInt, len(rrInt))
	fmt.Println(rrInt)
}

func isValidAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	rrS1 := []rune(s1)
	insertionSort(rrS1, len(s1))
	rrS2 := []rune(s2)
	insertionSort(rrS2, len(s2))

	if string(rrS1) != string(rrS2) {
		return false
	}

	return true
}

func insertionSort(rr []rune, n int) {

	if n <= 1 {
		return
	}

	insertionSort(rr, n-1)

	last := rr[n-1]
	counter := n - 2

	for counter >= 0 && rr[counter] > last {
		rr[counter+1] = rr[counter]
		counter--
	}
	rr[counter+1] = last
}

func insertionSortInt(rr []int, n int) {

	if n <= 1 {
		return
	}

	insertionSortInt(rr, n-1)

	last := rr[n-1]
	counter := n - 2

	for counter >= 0 && rr[counter] > last {
		rr[counter+1] = rr[counter]
		counter--
	}
	rr[counter+1] = last
}
