package main

import "fmt"

// source: https://leetcode.com/problems/add-two-numbers/
func main() {
	tcc := []struct {
		l1 *ListNode
		l2 *ListNode
		l3 *ListNode
	}{
		{
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l3: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 6,
					},
				},
			},
		},
		{
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			l3: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
		{
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l3: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
		{
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l3: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 7,
					},
				},
			},
		},
		{
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 7,
					},
				},
			},
			l3: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
		},
		{
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 8,
				},
			},
			l2: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 7,
					},
				},
			},
			l3: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			l1: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
				},
			},
			l2: &ListNode{
				Val: 1,
			},
			l3: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	for _, tc := range tcc {
		node := addTwoNumbers(tc.l1, tc.l2)

		res := tc.l3
		isFailed := false
		for res != nil {
			if res != nil && node == nil {
				isFailed = true
				break
			}

			if node.Val == res.Val {
				node = node.Next
				res = res.Next
			} else {
				isFailed = true
				break
			}
		}

		if isFailed {
			fmt.Println("Test Failed")
			continue
		}
		fmt.Println("Test successfull")
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	calcTogether(l1, l2, 0)
	return l1
}

func calcTogether(l1 *ListNode, l2 *ListNode, c int) {
	if l1 == nil && l2 == nil {
		return
	}

	var (
		carry, sum int = 0, c
	)
	if l1 != nil {
		sum += l1.Val
	}
	if l2 != nil {
		sum += l2.Val
	}

	l1.Val = sum % 10
	carry = sum / 10

	if l2 == nil {
		if l1.Next == nil && carry != 0 {
			l1.Next = &ListNode{Val: carry}
			return
		}

		calcTogether(l1.Next, nil, carry)
		return
	}
	if l1.Next == nil {
		if l2.Next != nil {
			l1.Next = l2.Next
			calcTogether(l1.Next, nil, carry)
			return
		}

		if l2.Next == nil && carry != 0 {
			l1.Next = &ListNode{Val: carry}
			return
		}
	}

	calcTogether(l1.Next, l2.Next, carry)
}
