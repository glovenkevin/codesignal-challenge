package main

import (
	"fmt"
	"math"
)

func main() {
	param := []int32{
		5, 4, 4, 2, 2, 8,
	}
	fmt.Println(cutTheSticks(param))
}

func cutTheSticks(arr []int32) []int32 {
	memo := []int32{}
	return cutSticks(arr, memo)
}

func cutSticks(arr []int32, memo []int32) []int32 {

	memo = append(memo, int32(len(arr)))

	var min int32 = math.MaxInt32
	for _, ar := range arr {
		if ar < min {
			min = ar
		}
	}

	sizeArr := len(arr)
	for i := 0; i < sizeArr; i++ {
		if arr[i]-min > 0 {
			arr[i] -= min
		} else {
			arr = append(arr[:i], arr[i+1:]...)
			sizeArr--
			i--
		}
	}

	if len(arr) == 0 {
		return memo
	}

	return cutSticks(arr, memo)
}
