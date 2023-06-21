package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// source: https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list
func main() {
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var nodeLength int
	node := head
	for node != nil {
		nodeLength++
		node = node.Next
	}

	node = head
	for i := 0; i < nodeLength/2; i++ {
		if i+1 == nodeLength/2 {
			node.Next = node.Next.Next
			break
		}
		node = node.Next
	}
	return head
}

/*
algorithm: walk twice as fast of the node, when the fast node reach end,
the slow one absolutely positioned on middle
*/
func deleteMiddleSimple(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow := head
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return head
}
