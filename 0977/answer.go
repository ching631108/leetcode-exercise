package main

import "fmt"

func sortedSquares(A []int) []int {
	// 生成负数结果集
	result1 := []int{}
	// 生成非负数结果集
	result2 := []int{}
	// 计算平方集合
	for _, item := range A {
		tmpResult := item * item
		if item < 0 {
			result1 = append(result1, tmpResult)
		} else {
			result2 = append(result2, tmpResult)
		}
	}
	// 生成正序的负数结果集
	result11 := []int{}
	for i := (len(result1) - 1); i >= 0; i-- {
		result11 = append(result11, result1[i])
	}
	// 拼接返回值
	returnData := []int{}
	tmp1Index := 0
	tmp2Index := 0
	for i := 0; i < len(A); i++ {
		if tmp1Index == len(result11) {
			for _, item := range result2[tmp2Index:] {
				returnData = append(returnData, item)
			}
			break
		}
		if tmp2Index == len(result2) {
			for _, item := range result11[tmp1Index:] {
				returnData = append(returnData, item)
			}
			break
		}
		fmt.Println("ready to append")
		if result11[tmp1Index] < result2[tmp2Index] {
			returnData = append(returnData, result11[tmp1Index])
			fmt.Println(len(returnData))
			tmp1Index++
		} else {
			returnData = append(returnData, result2[tmp2Index])
			fmt.Println(len(returnData))
			tmp2Index++
		}
	}
	return returnData[:]
}
