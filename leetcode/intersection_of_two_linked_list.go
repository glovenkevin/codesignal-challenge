package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/intersection-of-two-linked-lists/
// note: Could you write a solution that runs in O(m + n) time and use only O(1) memory?
func main() {
	itc := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}

	l1 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  1,
			Next: itc,
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  1,
				Next: itc,
			},
		},
	}

	fmt.Println("Result: ", getIntersectionNode2(l1, l2))
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lengthA, lengthB int32 = 1, 1
	var dumA, dumB *ListNode = headA, headB
	for dumA.Next != nil || dumB.Next != nil {
		if dumA.Next != nil {
			dumA = dumA.Next
			lengthA++
		}
		if dumB.Next != nil {
			dumB = dumB.Next
			lengthB++
		}
	}
	if dumA != dumB {
		return nil
	}

	for lengthA != lengthB {
		if lengthA > lengthB {
			headA = headA.Next
			lengthA--
			continue
		}
		if lengthA < lengthB {
			headB = headB.Next
			lengthB--
			continue
		}
	}

	for headA != nil {
		if headA != headB {
			headA = headA.Next
			headB = headB.Next
			continue
		}
		return headA
	}

	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB
	for l1 != l2 {
		if l1 == nil {
			l1 = headB
		} else {
			l1 = l1.Next
		}

		if l2 == nil {
			l2 = headA
		} else {
			l2 = l2.Next
		}
	}

	return l1
}
