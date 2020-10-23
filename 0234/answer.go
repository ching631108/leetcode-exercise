package main

func isPalindrome(head *ListNode) bool {
	// 处理链表的边界条件
	// 空元素
	if head == nil {
		return true
	}
	// 仅有一条数据
	if head.Next == nil {
		return true
	}
	// 仅有两条数据
	if head.Next.Next == nil {
		if head.Next.Val == head.Val {
			return true
		}
		return false
	}
	// 以上三种情况都会导致快指针寸步难行，手动处理
	// 快慢指针寻找链表中点
	faster := head
	slower := head
	for faster != nil && faster.Next != nil {
		faster = faster.Next.Next
		slower = slower.Next
	}
	// 从慢指针开始，将慢指针指向的链表反转
	tmpLastHead := slower
	slower = slower.Next
	tmpLastHead.Next = nil
	tmpNextHead := slower.Next
	for {
		slower.Next = tmpLastHead
		tmpLastHead = slower
		if tmpNextHead == nil {
			break
		}
		slower = tmpNextHead
		tmpNextHead = tmpNextHead.Next
	}
	for slower != nil && head != nil {
		if slower.Val != head.Val {
			return false
		}
		slower = slower.Next
		head = head.Next
	}
	return true
}
