package main

import "fmt"

func main() {
	fmt.Println(lineUp("LLARL")) // result: 3
}

func lineUp(commands string) int {
	count := 0 // 2
	lineUp := true
	commandBefore := ""                     // L
	for _, char := range []rune(commands) { // L
		currentCommand := string(char)
		if lineUp && currentCommand == "A" {
			count++
		} else {
			//
			if commandBefore != "" && !lineUp && currentCommand != "A" {
				lineUp = true
				count++
			} else {
				lineUp = false
			}
			commandBefore = currentCommand
			// lineup false
		}
	}
	return count
}
