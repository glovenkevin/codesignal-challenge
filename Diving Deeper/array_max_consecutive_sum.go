package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 4}
	fmt.Println(arrayMaxConsecutiveSum(arr, 3)) // 9
}

func arrayMaxConsecutiveSum(inputArray []int, k int) int {
    largest := 0
    counter := 0
    for counter < len(inputArray)-(k-1) {
        var tmp int
        for i:=counter; i<counter+k; i++ {
            tmp += inputArray[i]
        }
        if tmp > largest {
            largest = tmp
        }
        counter++
    }        
    return largest
}
