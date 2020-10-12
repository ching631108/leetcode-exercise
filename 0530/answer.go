package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	// 初始化问题答案
	answer := math.MaxInt64
	// 初始化上一个节点的缓存
	last := -1
	viewTree(root, &last, &answer)
	return answer
}

//
// 中序遍历一棵树
//
func viewTree(root *TreeNode, last *int, answer *int) {
	if root.Left != nil {
		viewTree(root.Left, last, answer)
	}
	// 判断上一个节点和本结点之间的差值
	if -1 != *last {
		diff := root.Val - *last
		if diff < *answer {
			*answer = diff
		}
	}
	// 处理结束，更新上一个节点的值
	*last = root.Val
	if root.Right != nil {
		viewTree(root.Right, last, answer)
	}
}
