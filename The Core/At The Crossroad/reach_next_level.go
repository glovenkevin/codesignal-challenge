package main

import "fmt"

func main() {
	fmt.Println(reachNextLevel(10, 15, 5)) // expected: true
}

func reachNextLevel(experience int, threshold int, reward int) bool {
    return experience + reward >=  threshold
}
