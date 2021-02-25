package main

import (
	"fmt"
)

func main() {
	fmt.Println(differentSymbolsNaive("cabca")) // return 3 different
}

func differentSymbolsNaive(s string) int {
    flag := make(map[rune]bool)
    for _, val := range []rune(s) {
        if _, exist := flag[val]; !exist {
            flag[val] = true
        }
    }
    return len(flag)
}
