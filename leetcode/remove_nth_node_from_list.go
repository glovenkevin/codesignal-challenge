package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/remove-nth-node-from-end-of-list/
func main() {
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	traceToEnd(head, &n)
	if n == 0 {
		return head.Next
	}
	return head
}

func traceToEnd(head *ListNode, n *int) *ListNode {
	if head == nil {
		return nil
	}

	node := traceToEnd(head.Next, n)

	if *n == 0 {
		head.Next = nil
		if node != nil {
			head.Next = node.Next
		}
	}
	*n--
	return head
}
