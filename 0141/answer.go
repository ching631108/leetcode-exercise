package main

//
// 方法一，使用 map 的 key 作为缓存，时间复杂度 O(n) 空间复杂度 O(n)
//
func hasCycleWithMap(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil {
		if m[head] == true {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}

//
// 方法2, 不使用 map 作为缓存，使用快慢指针法
//
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
