package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var head *ListNode
	list := []int{1, 2, 3, 4, 5}
	for i, _ := range list {
		push(&head, list[len(list)-1-i])
	}
	display(head)

	node := reverseKGroup(head, 3)

	display(node)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func length(head *ListNode) int {
	var l int
	curr := head
	for curr != nil {
		curr = curr.Next
		l++
	}

	return l
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	var pre, next *ListNode
	var count int

	length := length(head)
	if length < k {
		return head
	}

	for curr != nil && count < k {
		next = curr.Next
		curr.Next = pre
		pre = curr
		curr = next
		count++
	}

	if next != nil {
		head.Next = reverseKGroup(next, k)
	}

	return pre
}

func push(head **ListNode, val int) {
	node := &ListNode{
		Val:  val,
		Next: *head,
	}

	*head = node
}

func display(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val)
		curr = curr.Next
	}

	fmt.Print("\n")
}
