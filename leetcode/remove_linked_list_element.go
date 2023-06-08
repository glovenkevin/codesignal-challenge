package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/remove-linked-list-elements/
func main() {
	tc := []struct {
		head *ListNode
		val  int
		res  *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: &ListNode{Val: 6},
								},
							},
						},
					},
				},
			},
			val: 6,
			res: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
	}
test_loop:
	for _, t := range tc {
		res := removeElements(t.head, t.val)

		for res != nil {
			if res == nil || t.res == nil {
				fmt.Println("test failed")
				continue test_loop
			}
			if res.Val != t.res.Val {
				fmt.Println("test failed")
				continue test_loop
			}

			res = res.Next
			t.res = t.res.Next
		}

		fmt.Println("test success")
	}
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	node := removeElements(head.Next, val)
	if head.Val == val {
		return node
	}

	head.Next = node
	return head
}
