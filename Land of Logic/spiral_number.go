package main

import "fmt"

func main() {
	n := 3
	fmt.Println(spiralNumbers(n))
}

func spiralNumbers(n int) [][]int {
	total := n * n
	idx := 1
	rtn := make([][]int, n)
	for i := 0; i < n; i++ {
		rtn[i] = append(make([]int, n))
	}

	x := 0
	y := 0
	originDif := n
	dif := originDif
	way := "r" // r: right, l: left, u: up, d: down
	change := 0
	first := true
	for total > 0 {

		rtn[x][y] = idx

		if dif == 1 {

			if first {
				first = false
				originDif--
			} else if change > 0 {
				change = 0
				originDif--
			} else {
				change++
			}

			dif = originDif

			if way == "r" {
				way = "d"
			} else if way == "d" {
				way = "l"
			} else if way == "l" {
				way = "u"
			} else if way == "u" {
				way = "r"
			}

		} else {
			dif--
		}

		if way == "r" {
			y++
		} else if way == "d" {
			x++
		} else if way == "l" {
			y--
		} else {
			x--
		}

		idx++
		total--
	}
	return rtn
}
