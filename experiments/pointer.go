package main

import "fmt"

func main() {

	// purpose: check if the initial slice is being reslice, what will happen with the memory allocation
	// we try to prove that even the slice got resliced, it's still point to the same reference
	c := []byte("akus")
	fmt.Println("c", c)
	fmt.Println("c0", &c[0])
	fmt.Println("c1", &c[1])
	fmt.Println("c2", &c[2])

	d := c[:1]
	fmt.Println("d0", &d[0])
	e := c[1:]
	fmt.Println("e", e)
	fmt.Println("e1", &e[0])
	fmt.Println("e2", &e[1])

	fmt.Println("")
	f := make([]byte, len(c))
	copy(f, c)
	e[1] = 'd'
	fmt.Println("*c", c)
	fmt.Println("*e", e)
	fmt.Println("*f", f)
}
