package main

import "fmt"

func main() {
	param := []int32{
		0, 0, 1, 0, 0, 1, 0,
	}
	fmt.Println(jumpingOnClouds(param)) // expected: 4
}

func jumpingOnClouds(cc []int32) int32 {

	var (
		sizeCc       = len(cc) - 1
		steps  int32 = 0
		index  int   = 0
	)

	for index != sizeCc {

		if index+2 <= sizeCc && cc[index+2] != 1 {
			steps += 1
			index += 2
			continue
		}

		if cc[index+1] != 1 {
			steps += 1
			index += 1
		}
	}
	return steps
}

func jumpingOnClouds2(cc []int32) int32 {

	var (
		count int32 = 0
		size  int   = len(cc)
	)
	for index := 0; index < size; index++ {
		if cc[index] == 0 {
			index++
		}
		count++
	}
	return count
}
