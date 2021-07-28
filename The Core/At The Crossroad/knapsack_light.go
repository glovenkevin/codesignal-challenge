package main

import "fmt"

func main() {
	fmt.Println(knapsackLight(10, 5, 15, 6, 5)) // expected: 10
}

func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
    weight := 0
    if weight1 + weight2 <= maxW {
        weight = value1 + value2
    } else if weight1 > maxW && weight2 > maxW {
        
    } else {
        if (value1 > value2 && weight1 <= maxW) || (weight2 > maxW) {
            weight = value1
        } else if (weight2 <= maxW) {
            weight = value2
        }
    }
    return weight
}
