package main

import "fmt"

func main() {
    data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    k := 3

    fmt.Println(extractEachKth(data, k))
}

func extractEachKth(inputArray []int, k int) []int {
    res := []int{}
    if k==1 {
        return res
    }
    counter := 1
    for i, val := range inputArray {
        if i!=(k*counter)-1 {
            res = append(res, val)
        } else {
            counter+=1
        }
    }
    return res
}
