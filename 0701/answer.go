package main

// 插入指定元素, 使用递归实现

func insertIntoBSTForRecursion(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root := &TreeNode{val, nil, nil}
		return root
	}
	if val > root.Val {
		if root.Right == nil {
			tmpNode := TreeNode{val, nil, nil}
			root.Right = &tmpNode
		} else {
			insertIntoBSTForRecursion(root.Right, val)
		}
	} else {
		if root.Left == nil {
			tmpNode := TreeNode{val, nil, nil}
			root.Left = &tmpNode
		} else {
			insertIntoBSTForRecursion(root.Left, val)
		}
	}
	return root
}

// 插入指定元素, 不使用递归实现

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root := &TreeNode{val, nil, nil}
		return root
	}
	var lastHeader **TreeNode
	tmpHeader := root
	for tmpHeader != nil {
		if val > tmpHeader.Val {
			lastHeader = &tmpHeader.Right
			tmpHeader = tmpHeader.Right
		} else {
			lastHeader = &tmpHeader.Left
			tmpHeader = tmpHeader.Left
		}
	}
	tmpHeader = &TreeNode{val, nil, nil}
	*lastHeader = tmpHeader
	return root
}
