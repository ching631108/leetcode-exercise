package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{1, nil, nil}
	initNodeData(&root)
	// printInitNode(&root)
	result := postorderTraversal(&root)
	fmt.Println(result)
}
func initNodeData(root *TreeNode) {

	node2 := TreeNode{2, nil, nil}
	root.Left = &node2

	node3 := TreeNode{3, nil, nil}
	root.Right = &node3

	node4 := TreeNode{4, nil, nil}
	node2.Left = &node4

	node5 := TreeNode{5, nil, nil}
	node2.Right = &node5

	node7 := TreeNode{7, nil, nil}
	node3.Right = &node7
}
func printInitNode(root *TreeNode) {
	if root != nil {
		fmt.Println(root.Val)
	}
	if root.Left != nil {
		printInitNode(root.Left)
	}
	if root.Right != nil {
		printInitNode(root.Right)
	}
}
