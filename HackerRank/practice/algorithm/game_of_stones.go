package main

import "fmt"

func main() {
	fmt.Println(gameOfStones(8)) // expected: Second
	fmt.Println(gameOfStones(1)) // expected: First
}

/*
	1 -> as first player, 2 -> as second player.
	Rules: player can only remove 2,3, or 5 stones each step.
		First game will always be started by first player.

	if the remain is 1 or 7 then it can be presumed that the Second player is the winner.
*/
func gameOfStones(n int32) string {
	s := n % 7
	if s == 0 || s == 1 {
		return "Second"
	}

	return "First"
}
