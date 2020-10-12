package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{236, nil, nil}
	initNodeData(root)
	// printInitNode(root)
	result := getMinimumDifference(root)
	fmt.Println(result)
}
func initNodeData(root *TreeNode) {

	node2 := TreeNode{104, nil, nil}
	root.Left = &node2

	node7 := TreeNode{701, nil, nil}
	root.Right = &node7

	node1 := TreeNode{911, nil, nil}
	node7.Right = &node1

	node3 := TreeNode{227, nil, nil}
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
