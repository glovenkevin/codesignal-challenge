package main

import "fmt"

func main() {

	p1 := []int{3, 2, 1}
	p2 := []int{1, 2, 3}
	fmt.Println(compareTriplets(p1, p2))

	p1 = []int{5, 6, 7}
	p2 = []int{3, 6, 10}
	fmt.Println(compareTriplets(p1, p2))

	p1 = []int{17, 28, 30}
	p2 = []int{99, 16, 8}
	fmt.Println(compareTriplets(p1, p2))

}

// a > b -> alice +1
// a < b -> bob +1
// a == b -> neither both get points
func compareTriplets(a []int, b []int) []int {
	res := make([]int, 2)
	for i := 0; i < 3; i++ {
		if a[i] == b[i] {
			continue
		}
		if a[i] < b[i] {
			res[1]++
			continue
		}

		if a[i] > b[i] {
			res[0]++
		}
	}
	return res
}
