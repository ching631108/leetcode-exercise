package main

func swapPairs(head *ListNode) *ListNode {
	fakeHead := &ListNode{0, head}
	perNode := fakeHead
	for head != nil {
		// 如果当前节点需要切换
		if head.Next != nil {
			perNode.Next = head.Next
			head.Next = head.Next.Next
			perNode.Next.Next = head
		}
		perNode = head
		head = head.Next
	}
	return fakeHead.Next
}
