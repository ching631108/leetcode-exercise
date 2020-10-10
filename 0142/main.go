package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	headHead := &ListNode{0, nil}
	initList(headHead)
	// printList(headHead)
	result := detectCycle(headHead)
	fmt.Println(result)
}

func initList(head *ListNode) {
	node1 := &ListNode{1, nil}
	head.Next = node1

	node2 := &ListNode{2, nil}
	node1.Next = node2

	node3 := &ListNode{3, nil}
	node2.Next = node3

	node4 := &ListNode{4, nil}
	node3.Next = node4

	node5 := &ListNode{5, nil}
	node4.Next = node5

	node6 := &ListNode{6, nil}
	node5.Next = node6
	node6.Next = node1
}
func printList(head *ListNode) {
	count := 0
	for head != nil && count < 100 {
		count = count + 1
		fmt.Println(head.Val, head.Next)
		head = head.Next
	}
}
