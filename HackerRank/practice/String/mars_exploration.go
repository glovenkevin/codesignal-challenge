package main

import (
	"fmt"
)

func main() {
	fmt.Println(marsExploration("SOSSPSSQSSOR"))
}

func marsExploration(s string) int32 {

	const SOS string = "SOS"
	diff := 0
	for i := 0; i < len(s); i += 3 {
		sos := s[i : i+3]
		if sos[0] != SOS[0] {
			diff++
		}
		if sos[1] != SOS[1] {
			diff++
		}
		if sos[2] != SOS[2] {
			diff++
		}
	}

	return int32(diff)

}
