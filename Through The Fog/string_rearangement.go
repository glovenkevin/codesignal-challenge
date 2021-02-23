package main

import (
	"fmt"
	"sort"
)

func main() {
	tes1 := []string{"aba", "bbb", "bab"}
	fmt.Println(stringsRearrangement(tes1))

	tes2 := []string{"ab", "bb", "aa"}
	fmt.Println(stringsRearrangement(tes2))

	tes3 := []string{"ab", "ad", "ef", "eg"}
	fmt.Println(stringsRearrangement(tes3))

	tes4 := []string{"abc", "abx", "axx", "abx", "abc"}
	fmt.Println(stringsRearrangement(tes4))

	tes5 := []string{"ff", "gf", "af", "ar", "hf"}
	fmt.Println(stringsRearrangement(tes5))

	data := []rune("abd")
	sum := 0
	for i:=0; i<len(data); i++ {
		sum += int(data[i])
	}
	fmt.Println(sum)
}

func checkArray(inputArray []string) bool {
    for i:=0; i<len(inputArray)-1; i++ {
        if !diffOneChar(inputArray[i], inputArray[i+1]) {
            return false
        }
    }
    
    return true
}

func diffOneChar(word1 string, word2 string) bool {
    count := 0
    for idx:=0; idx<len(word1); idx++ {
        if (word1[idx] != word2[idx]) {
            count += 1
        }
    }
    
    return count==1
}

func permutation(arr []string) [][]string {
    var helper func([]string, int)
    res := [][]string{}
    
    helper = func(arr []string, n int) {
        if n==1 {
            tmp := make([]string, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++{
                helper(arr, n - 1)
                if n % 2 == 1{
                    tmp := arr[i]
                    arr[i] = arr[n - 1]
                    arr[n - 1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n - 1]
                    arr[n - 1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
}

func stringsRearrangement(inputArray []string) bool {
    sort.Strings(inputArray)
    if (checkArray(inputArray)) {
        return true
    }
    
    listPerm := permutation(inputArray)
    for idx:=0; idx<len(listPerm); idx++ {
        if(checkArray(listPerm[idx])) {
            return true
        }
    }
    
    return false
}
