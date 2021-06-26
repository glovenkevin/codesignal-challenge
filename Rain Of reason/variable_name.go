package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("normal: ", variableName("t_123aa"))               // True
	fmt.Println("spesial char: ", variableName("t_123-aa"))        // False
	fmt.Println("satu char saja: ", variableName("t"))             // True
	fmt.Println("satu char spesial karakter: ", variableName("_")) // True
	fmt.Println("normal: ", variableName("__asd12"))               // True
}

func variableName(name string) bool {
	regex := regexp.MustCompile("^[a-zA-Z_][a-zA-Z0-9_]*$")
	return regex.MatchString(name)
}
