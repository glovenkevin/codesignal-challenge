package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/swap-nodes-in-pairs/
func main() {
	tc := []struct {
		head *ListNode
		res  *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			res: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		},
	}

test_loop:
	for _, t := range tc {
		res := swapPairs(t.head)

		for t.res != nil {
			if res != nil && res.Val == t.res.Val {
				t.res = t.res.Next
				res = res.Next
				continue
			}

			fmt.Println("test failed")
			continue test_loop
		}

		fmt.Println("test success")
	}
}

func swapPairs(head *ListNode) *ListNode {
	var (
		bef  *ListNode
		curr = head
	)
	for curr != nil && curr.Next != nil {
		second := curr.Next
		third := second.Next
		curr, second = second, curr
		curr.Next = second
		second.Next = third

		if bef == nil {
			head = curr
		}
		if bef != nil {
			bef.Next = curr
		}

		bef = second
		curr = third
	}

	return head
}
