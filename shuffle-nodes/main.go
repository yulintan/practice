// Given a linked list, uniformly shuffle the nodes. What if we want to prioritize space over time?
// 1 2 3 4 5 6 7 8 9 10
// would be
// 6 1 7 2 8 3 9 4 10 5
// assume count of list is even

package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var size int

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	size = len(a)

	list := buildList(a)
	printList(list)
	list = shuffle(list)
	printList(list)
	fmt.Println("Done")
}

func buildList(a []int) *Node {
	var head *Node
	for _, n := range a {
		head = addNode(head, Node{Value: n})
	}

	return head
}

func shuffle(head *Node) *Node {
	if size == 0 || size == 2 {
		return head
	}

	var slow, fast *Node
	slow = head
	fast = head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	fast = slow
	slow = head

	var tail *Node
	tail = fast
	for fast != nil {
		slowNext := slow.Next
		fastNext := fast.Next

		slow.Next = fastNext
		fast.Next = slow

		fast = fastNext
		slow = slowNext
	}

	return tail
}

func addNode(head *Node, node Node) *Node {
	if head == nil {
		head = &node

		return head
	}

	var curr *Node
	curr = head
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = &node

	return head
}

func printList(head *Node) {
	var curr = head
	for curr != nil {
		fmt.Print(curr.Value, " ")
		curr = curr.Next
	}

	fmt.Printf("\n")
}
