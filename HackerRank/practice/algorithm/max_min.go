package main

import (
	"fmt"
	"math"
)

func main() {
	param := []int32{1, 4, 3, 2}
	fmt.Println(maxMin(2, param)) // Expected: 1

	param2 := []int32{10, 100, 300, 200, 1000, 20, 30}
	fmt.Println(maxMin(3, param2)) // Expected: 20
}

func maxMin(k int32, arr []int32) int32 {
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	var (
		arrs                = len(arr)
		minUnfairness int32 = math.MaxInt32
	)
	for i := 0; i < arrs; i++ {
		r := i + int(k)
		if r > arrs {
			break
		}

		newArr := arr[i:r]
		unfairness := calcUnfairness(newArr)
		if minUnfairness > unfairness {
			minUnfairness = unfairness
		}
	}

	return minUnfairness
}

func calcUnfairness(arr []int32) int32 {
	var (
		min int32 = math.MaxInt32
		max int32 = 0
	)

	for i := 0; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
		if max < arr[i] {
			max = arr[i]
		}
	}

	return max - min
}

func mergeSort(aa []int32, l int, r int) {
	if l < r {
		mid := (l + r - 1) / 2
		mergeSort(aa, l, mid)
		mergeSort(aa, mid+1, r)
		merge(aa, l, mid, r)
	}
}

func merge(aa []int32, l int, m int, r int) {
	ll := make([]int32, len(aa[l:m+1]))
	copy(ll, aa[l:m+1])
	rr := make([]int32, len(aa[m+1:r+1]))
	copy(rr, aa[m+1:r+1])

	var (
		ls      int = len(ll)
		rs      int = len(rr)
		i       int = 0
		j       int = 0
		counter int = l
	)

	for i < ls || j < rs {
		if i == ls && j < rs {
			aa[counter] = rr[j]
			j++
			counter++
			continue
		}

		if i < ls && j == rs {
			aa[counter] = ll[i]
			i++
			counter++
			continue
		}

		if ll[i] <= rr[j] {
			aa[counter] = ll[i]
			i++
			counter++
		} else {
			aa[counter] = rr[j]
			j++
			counter++
		}
	}
}
