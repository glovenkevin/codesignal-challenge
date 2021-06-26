package main

import "fmt"

func main() {
	fmt.Println(alphabeticShift("crazy"))
}

func alphabeticShift(inputString string) string {
	listChar := []rune(inputString)
	rtn := ""
	for i := 0; i < len(listChar); i++ {
		if "Z" == string(listChar[i]) || "z" == string(listChar[i]) {
			rtn += string(listChar[i] - 25)
		} else {
			rtn += string(listChar[i] + 1)
		}
	}
	return rtn
}
