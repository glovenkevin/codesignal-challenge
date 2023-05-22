package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/linked-list-cycle/
func main() {
}

// this methods was so called fast and slow approach.
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast := head
	slow := head
	for slow.Next != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}

	return false
}
