package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/reverse-linked-list/
func main() {
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newStart := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newStart
}
