package main

import "fmt"

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	initHead := &ListNode{1, nil}
	initList(initHead, []int{2})
	reorderList(initHead)
	printInitList(initHead)
}

//
// 初始化列表
//
func initList(head *ListNode, list []int) {
	for _, item := range list {
		tmpNode := &ListNode{item, nil}
		head.Next = tmpNode
		head = tmpNode
	}
}

//
// 打印链表
//
func printInitList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
