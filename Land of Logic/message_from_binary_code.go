package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	param := "010010000110010101101100011011000110111100100001" // expected: Hello!
	fmt.Println(messageFromBinaryCode(param))
}

func messageFromBinaryCode(code string) string {
	countSubstring := 0
	loopCounter := len(code) / 8
	decryptWord := ""
	for loopCounter > 0 {
		bit := code[countSubstring : countSubstring+8]
		decryptWord += convert8BitToAscii(bit)
		countSubstring += 8
		loopCounter--
	}
	return decryptWord
}

func convert8BitToAscii(bit string) string {
	arrBit := []rune(bit)
	binaryPow := 7.0
	ascii := 0.0
	for _, val := range arrBit {
		binary, _ := strconv.Atoi(string(val))
		if binary != 0 {
			ascii += math.Pow(2, binaryPow)
		}
		binaryPow--
	}
	return string(int(ascii))
}
