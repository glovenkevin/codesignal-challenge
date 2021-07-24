package main

import (
	"fmt"
)

func main() {
	fmt.Println(electionsWinners([2, 3, 5, 2], 2))
}

func electionsWinners(votes []int, k int) int {
    
    var biggestVotes = 0
    var countBiggestVotes = 0
    
    for _, val := range votes {
        if (biggestVotes < val) {
            biggestVotes = val
            countBiggestVotes = 1
        } else if (biggestVotes == val) {
            countBiggestVotes++
        }
    }
    
    if countBiggestVotes > 1 && k == 0 {
        return 0
    }
    
    if k > 0 {
        var countBiggestWithNoVotes = 0
        for _, val := range votes {
            if (val+k > biggestVotes) {
                countBiggestWithNoVotes++
            }
        }
        return countBiggestWithNoVotes
    } else {
        return countBiggestVotes
    }
    
}
