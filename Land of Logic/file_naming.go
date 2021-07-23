package main

import (
	"fmt"
	"strconv"
)

func main() {
	param := []string{
		"doc",
		"doc",
		"image",
		"doc(1)",
		"doc",
	}
	fmt.Println(fileNaming(param))

	param2 := []string{
		"a(1)",
		"a(6)",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
		"a",
	}
	fmt.Println(fileNaming(param2))

	param3 := []string{
		"dd",
		"dd(1)",
		"dd(2)",
		"dd",
		"dd(1)",
		"dd(1)(2)",
		"dd(1)(1)",
		"dd",
		"dd(1)",
	}
	fmt.Println(fileNaming(param3))
}

func fileNaming(names []string) []string {
	countNames := make(map[string]int)
	arrNewNames := []string{}
	for _, docName := range names {
		count := countNames[docName]
		newNames := ""
		if count == 0 {
			newNames = docName
			countNames[docName]++
		} else {
			newNames = docName + "(" + strconv.Itoa(countNames[docName]) + ")"
			for countNames[newNames] != 0 {
				countNames[docName]++
				newNames = docName + "(" + strconv.Itoa(countNames[docName]) + ")"
			}
			countNames[docName]++
			countNames[newNames]++
		}
		arrNewNames = append(arrNewNames, newNames)
	}

	return arrNewNames
}
