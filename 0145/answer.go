package main

//
// Definition for a binary tree node.
// type TreeNode struct {
//     Val int
//     Left *TreeNode
//     Right *TreeNode
// }

func postorderTraversal(root *TreeNode) []int {
	returnData := []int{}
	if root != nil {
		_getData(root, &returnData)
	}
	return returnData
}

// 关于后续遍历的递归实现
func _getData(root *TreeNode, list *[]int) {
	if root.Left != nil {
		_getData(root.Left, list)
	}
	if root.Right != nil {
		_getData(root.Right, list)
	}
	*list = append(*list, root.Val)
}
