package main

func connectV1(root *Node) *Node {
	lastLineHead := root
	var tmpNextRight *Node
	for lastLineHead != nil {
		tmpCurrntLineHead := lastLineHead
		tmpNextRight = nil
		lastLineHead = nil
		for tmpCurrntLineHead != nil && tmpCurrntLineHead.Left != nil {
			if lastLineHead == nil {
				lastLineHead = tmpCurrntLineHead.Left
			}
			if tmpNextRight != nil {
				tmpNextRight.Next = tmpCurrntLineHead.Left
			}
			tmpCurrntLineHead.Left.Next = tmpCurrntLineHead.Right
			tmpNextRight = tmpCurrntLineHead.Right
			tmpCurrntLineHead = tmpCurrntLineHead.Next
		}
	}
	return root
}

// 换一个递归写法
func connect(root *Node) *Node {
	// 当前节点为空或者下层节点为空
	if root == nil || root.Left == nil {
		return root
	}
	root.Left.Next = root.Right
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}
	// 根据 leetcode 的结果，尾部递归比上述方法要快
	connect(root.Left)
	connect(root.Right)
	return root
}
