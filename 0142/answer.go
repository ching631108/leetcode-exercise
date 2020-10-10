package main

func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for head != nil {
		if m[head] {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}
