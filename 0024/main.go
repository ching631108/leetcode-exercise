package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// head := initListNode([]int{1, 2, 3, 4, 5, 6})
	// result := swapPairs(head)
	result := swapPairs(nil)
	printListNode(result)
}

//
// initListNode 初始化链表
//
func initListNode(list []int) *ListNode {
	fakeNode := &ListNode{0, nil}
	lastNode := fakeNode
	for _, item := range list {
		tmpNode := &ListNode{item, nil}
		lastNode.Next = tmpNode
		lastNode = tmpNode
	}
	return fakeNode.Next
}

//
//	打印链表
//
func printListNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
