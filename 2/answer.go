package main

//
// Definition for singly-linked list.
// type ListNode struct {
//     Val int
//     Next *ListNode
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Header := l1
	l2Header := l2
	var tmpHeader *ListNode
	var result *ListNode
	firstHeader := true
	plusRemainder := 0
	plusResult := 0
	for l1Header != nil || l2Header != nil || plusRemainder != 0 {
		tmp1 := 0
		if l1Header != nil {
			tmp1 = l1Header.Val
			l1Header = l1Header.Next
		}
		tmp2 := 0
		if l2Header != nil {
			tmp2 = l2Header.Val
			l2Header = l2Header.Next
		}

		plusResult, plusRemainder = plus(tmp1, tmp2, plusRemainder)
		tmpResult := ListNode{plusResult, nil}
		if firstHeader {
			result = &tmpResult
			firstHeader = false
		} else {
			tmpHeader.Next = &tmpResult
		}
		tmpHeader = &tmpResult
	}

	return result
}

/***
 * 就算两个数字的和
 * 返回个位和余数
 * @return
 * int 两个数字之和的个位
 * int 两个数字之和的十位
 */

func plus(a1 int, a2 int, a3 int) (int, int) {
	result := a1 + a2 + a3
	return result % 10, result / 10
}
