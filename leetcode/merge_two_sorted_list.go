package main

import "fmt"

// source: https://leetcode.com/problems/merge-two-sorted-lists/
func main() {
	pp := []struct {
		l1  []int
		l2  []int
		res []int
	}{
		{
			l1:  []int{1, 2, 4},
			l2:  []int{1, 3, 4},
			res: []int{1, 1, 2, 3, 4, 4},
		},
		{
			l1:  []int{},
			l2:  []int{},
			res: []int{},
		},
		{
			l1:  []int{},
			l2:  []int{0},
			res: []int{0},
		},
		{
			l1:  []int{-9, -7, -3, 1, 2},
			l2:  []int{-7, -7, -3, 2, 4},
			res: []int{-9, -7, -7, -7, -3, -3, 1, 2, 2, 4},
		},
		{
			l1:  []int{2},
			l2:  []int{1},
			res: []int{1, 2},
		},
	}

	for i, p := range pp {
		l1 := CreateLinkedList(p.l1)
		l2 := CreateLinkedList(p.l2)

		res := mergeTwoLists(l1, l2)
		var rr []int
		for res != nil {
			rr = append(rr, res.Val)
			res = res.Next
		}

		fmt.Printf("%d. Result: %v -> Expected: %v \n", i+1, rr, p.res)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(pp []int) *ListNode {
	var init, tail *ListNode
	for _, p := range pp {
		if init == nil {
			init = &ListNode{Val: p}
			tail = init
			continue
		}
		n := &ListNode{Val: p}
		tail.Next = n
		tail = n
	}
	return init
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var tail ListNode
	merge(list1, list2, &tail)
	return tail.Next
}

func merge(l1 *ListNode, l2 *ListNode, tail *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil && l2 != nil {
		tail.Next = l2
		return nil
	}
	if l1 != nil && l2 == nil {
		tail.Next = l1
		return nil
	}

	if tail.Val == 0 {
		if l1.Val >= l2.Val {
			tail.Next = l2
			tail = tail.Next
			l2 = l2.Next
			return merge(l1, l2, tail)
		}
		if l1.Val <= l2.Val {
			tail.Next = l1
			tail = tail.Next
			l1 = l1.Next
			return merge(l1, l2, tail)
		}
	}

	if l1.Val >= tail.Val && l1.Val <= l2.Val {
		tail.Next = l1
		tail = tail.Next
		return merge(l1.Next, l2, tail)
	}

	if l2.Val >= tail.Val && l2.Val <= l1.Val {
		tail.Next = l2
		tail = tail.Next
		return merge(l1, l2.Next, tail)
	}

	return merge(l1, l2, tail)
}
