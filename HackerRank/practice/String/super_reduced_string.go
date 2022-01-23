package main

import "fmt"

func main() {
	fmt.Println(superReducedString("aaabccddd")) // Expected: abd
}

/*
	Another way to solve this is to replace all char that
	match pattern regex /(\w)\1/g
*/

func superReducedString(s string) string {

	rr := []rune(s)
	rtn := ""
	size := len(rr) - 1
	for i := 0; i <= size; i++ {
		if i == size || rr[i] != rr[i+1] {
			rtn += string(rr[i])
			continue
		}

		if rr[i] == rr[i+1] {
			i++
		}
	}

	if s != rtn {
		return superReducedString(rtn)
	}

	if rtn == "" {
		return "Empty String"
	}

	return rtn
}
