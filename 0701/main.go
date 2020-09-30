package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{4, nil, nil}
	initNodeData(&root)
	insertIntoBST(&root, 5)
	printInitNode(&root)
}
func initNodeData(root *TreeNode) {

	node2 := TreeNode{2, nil, nil}
	root.Left = &node2

	node7 := TreeNode{7, nil, nil}
	root.Right = &node7

	node1 := TreeNode{1, nil, nil}
	node2.Left = &node1

	node3 := TreeNode{3, nil, nil}
	node2.Right = &node3

}
func printInitNode(root *TreeNode) {
	if root.Left != nil {
		printInitNode(root.Left)
	}
	if root != nil {
		fmt.Println(root.Val)
	}
	if root.Right != nil {
		printInitNode(root.Right)
	}

}
