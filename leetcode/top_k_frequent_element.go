package main

import "fmt"

// source: https://leetcode.com/problems/top-k-frequent-elements/
func main() {
	tc := []struct {
		nums []int
		k    int
		res  []int
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			res:  []int{1, 2},
		},
		{
			nums: []int{1},
			k:    1,
			res:  []int{1},
		},
	}

test_loop:
	for _, t := range tc {
		res := topKFrequent(t.nums, t.k)

		for _, r := range t.res {
			notExist := true
			for _, h := range res {
				if r == h {
					notExist = false
					break
				}
			}

			if notExist {
				fmt.Println("test failed")
				continue test_loop
			}
		}

		fmt.Println("test success")
	}
}

func topKFrequent(nums []int, k int) []int {
	var (
		totalNumMap = make(map[int]int, len(nums))
		freq        = make([][]int, len(nums)+1)
		ret         = make([]int, 0, len(nums))
	)
	for _, n := range nums {
		totalNumMap[n]++
	}

	for k, v := range totalNumMap {
		freq[v] = append(freq[v], k)
	}

	for i := len(freq) - 1; i >= 0; i-- {
		if freq[i] != nil {
			ret = append(ret, freq[i]...)
		}

		if len(ret) >= k {
			return ret[:k]
		}
	}

	return ret
}
