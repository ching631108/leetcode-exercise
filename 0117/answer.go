package main

//
// Definition for a Node.
// type Node struct {
//     Val int
//     Left *Node
//     Right *Node
//     Next *Node
// }
//

func connect01(root *Node) *Node {
	// 生成一个用户存储同级别节点的切片
	if root == nil {
		return root
	}
	nodeArr := []*Node{root}
	for {
		// 清空subArr
		subArr := []*Node{}
		// 遍历父级节点, 寻找子节点
		for _, currentNode := range nodeArr {
			if currentNode.Left != nil {
				subArr = append(subArr, currentNode.Left)
			}
			if currentNode.Right != nil {
				subArr = append(subArr, currentNode.Right)
			}
		}
		// 更新子节点的 next 标签
		subArrLength := len(subArr)
		for index, currentSonNode := range subArr {
			if index < subArrLength-1 {
				currentSonNode.Next = subArr[index+1]
			} else {
				currentSonNode.Next = nil
			}
		}
		if len(subArr) == 0 {
			break
		}
		// 将子节点列表置为父节点列表
		nodeArr = make([]*Node, len(subArr))
		copy(nodeArr, subArr)
	}
	return root
}

//
// 根据官方题解提示 如果已经建立了下一层的 next 指针
// 则下一层实际上形成了一个链表
// 可以使用链表构建相关数据而不使用数组
//
func connect(root *Node) *Node {
	// 生成一个用户存储同级别节点的切片
	if root == nil {
		return root
	}
	currentLineHeader := root
	for currentLineHeader != nil {
		tmpHeader := currentLineHeader
		currentLineHeader = nil
		// 遍历本行链表
		var nextLineHeader *Node
		for tmpHeader != nil {
			if tmpHeader.Left != nil {
				fixData(&currentLineHeader, &nextLineHeader, tmpHeader.Left)
			}
			if tmpHeader.Right != nil {
				fixData(&currentLineHeader, &nextLineHeader, tmpHeader.Right)
			}
			tmpHeader = tmpHeader.Next
		}
	}
	return root
}

func fixData(currentLineHeader **Node, nextLineHeader **Node, targetNode *Node) {
	if *currentLineHeader == nil {
		*currentLineHeader = targetNode
	}
	if *nextLineHeader == nil {
		*nextLineHeader = targetNode
	} else {
		(*nextLineHeader).Next = targetNode
	}
	*nextLineHeader = targetNode
}
