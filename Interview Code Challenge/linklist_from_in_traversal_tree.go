package main

import "fmt"

type node struct {
	childLeft  *node
	childRight *node
	value      int
	// father *node
}

func main() {

	/*
	   Old Content below(Java):


	   // 1
	   // 2 3
	   // 4 nil 5 6

	   // 4 -> 2 -> 1 -> 5 -> 3 -> 6
	*/

	var (
		root, child1, child2, child3, child4, child5 = node{}, node{}, node{}, node{}, node{}, node{}

		child6 = node{}
	)

	root.value = 1
	child1.value = 2
	child2.value = 3
	child3.value = 4
	child4.value = 5
	child5.value = 6
	child6.value = 7

	root.childLeft = &child1
	root.childRight = &child2

	child1.childLeft = &child3
	child1.childRight = nil

	child2.childLeft = &child4
	child2.childRight = &child5

	child3.childRight = &child6

	// printInTraversalTree(&root)
	createLinkedList(&root, nil)

	var tmpNode *node = &child6
	var count = 0
	for count < 7 {
		count++
		fmt.Print(tmpNode.value)
		tmpNode = tmpNode.childRight
	}

}

func printInTraversalTree(root *node) {

	if root.childLeft != nil {
		printInTraversalTree(root.childLeft)
	}

	fmt.Print(root.value)

	if root.childRight != nil {
		printInTraversalTree(root.childRight)
	}
	return
}

func createLinkedList(root *node, currNode *node) *node {
	var (
		lastNode *node
	)

	if root.childLeft != nil {
		lastNode = createLinkedList(root.childLeft, currNode)
	}

	if lastNode != nil {
		lastNode.childRight = root
		lastNode = root
	}

	if currNode != nil && lastNode == nil {
		currNode.childRight = root
		lastNode = root
	}

	if root.childRight != nil {
		createLinkedList(root.childRight, lastNode)
	}

	if currNode != nil {
		currNode.childRight = root
	}

	lastNode = root
	return lastNode
}
