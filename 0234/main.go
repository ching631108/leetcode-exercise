package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}
	initList(head, []int{1, 3, 1})
	// printInitList(head.Next)
	result := isPalindrome(head.Next)
	fmt.Println(result)
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
