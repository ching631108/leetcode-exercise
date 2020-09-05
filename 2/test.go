package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	initData1 := ListNode{}
	initData1.initList([]int{0})
	// initData1.printList()
	initData2 := ListNode{}
	initData2.initList([]int{1, 2, 3, 4, 5})
	result := addTwoNumbers(&initData1, &initData2)
	result.printList()
	return
}

// 一些关于数据结构的辅助方法
func (header *ListNode) printList() {
	tmpHeader := header
	for {
		fmt.Println(tmpHeader.Val)
		if nil == tmpHeader.Next {
			return
		}
		tmpHeader = tmpHeader.Next
	}
}
func (header *ListNode) initList(listData []int) {
	firstHeader := true
	for _, v := range listData {
		tmpNode := ListNode{v, nil}
		if firstHeader {
			firstHeader = false
			header.Val = v
			continue
		}
		header.Next = &tmpNode
		header = &tmpNode
	}
	return
}
