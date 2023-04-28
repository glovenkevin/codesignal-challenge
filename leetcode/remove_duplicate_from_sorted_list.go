package main

import "fmt"

// source: https://leetcode.com/problems/remove-duplicates-from-sorted-list
func main() {
	tc := []struct {
		head *ListNode
		res  *ListNode
	}{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 3,
							},
						},
					},
				},
			},
			res: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}

test_loop:
	for _, t := range tc {
		n := deleteDuplicates(t.head)

		for n != nil && t.res != nil {
			if n.Val != t.res.Val {
				fmt.Println("test failed")
				continue test_loop
			}

			n = n.Next
			t.res = t.res.Next
		}
		fmt.Println("test success")
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	n := head
	for n != nil {
		if n.Next != nil && n.Val == n.Next.Val {
			n.Next = n.Next.Next
			continue
		}

		n = n.Next
	}

	return head
}
