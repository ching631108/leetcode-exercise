package main

// 使用快慢指针获取链表中间的位置
func reorderList(head *ListNode) {
	// 针对特殊的边缘情况进行处理
	if nil == head || // 传入空数据
		nil == head.Next || // 链表长度为 1
		nil == head.Next.Next { // 链表长度为 2
		return
	}
	fasterHead := head
	slowerHead := head
	for fasterHead != nil && fasterHead.Next != nil {
		slowerHead = slowerHead.Next
		fasterHead = fasterHead.Next.Next
	}
	// 从中间开始，将慢指针指向的链表反转
	lastHead := slowerHead
	slowerHead = slowerHead.Next
	lastHead.Next = nil
	for slowerHead != nil {
		// 暂存下一个节点
		tmpNext := slowerHead.Next
		slowerHead.Next = lastHead
		lastHead = slowerHead
		if tmpNext == nil {
			break
		}
		slowerHead = tmpNext
	}
	// 开始拼接交错链表
	for head != nil {
		tmpHead := head.Next
		head.Next = slowerHead
		if slowerHead == nil {
			break
		}
		slowerHead = slowerHead.Next
		head.Next.Next = tmpHead
		if head.Next == nil {
			break
		}
		head = head.Next.Next
	}
}
