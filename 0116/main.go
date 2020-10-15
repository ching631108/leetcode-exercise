package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {
	root := &Node{1, nil, nil, nil}
	initNodeData(root)
	connect(root)
	printInitNode(root)
}

func initNodeData(root *Node) {

	node2 := Node{2, nil, nil, nil}
	root.Left = &node2

	node3 := Node{3, nil, nil, nil}
	root.Right = &node3

	node4 := Node{4, nil, nil, nil}
	node2.Left = &node4

	node5 := Node{5, nil, nil, nil}
	node2.Right = &node5

	node6 := Node{6, nil, nil, nil}
	node3.Left = &node6

	node7 := Node{7, nil, nil, nil}
	node3.Right = &node7
}
func printInitNode(root *Node) {
	if root != nil {
		tmpStr := "null"
		if root.Next != nil {
			tmpStr = fmt.Sprintf("%d", root.Next.Val)
		}
		fmt.Println(root.Val, tmpStr)
	}
	if root.Left != nil {
		printInitNode(root.Left)
	}
	if root.Right != nil {
		printInitNode(root.Right)
	}
}
func printNodeSlice(list []*Node, tag string) {
	for _, currentSonNode := range list {
		tmpStr := "null"
		if currentSonNode.Next != nil {
			tmpStr = fmt.Sprintf("%d", currentSonNode.Next.Val)
		}
		fmt.Println(tag, currentSonNode.Val, tmpStr)
	}
}
