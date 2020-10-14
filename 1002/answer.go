package main

import (
	"math"
)

func commonChars(A []string) []string {
	result := []string{}
	var targetStr string
	// 寻找最短的字符串
	minLength := math.MaxInt64
	for _, str := range A {
		if len(str) < minLength {
			targetStr = str
			minLength = len(str)
		}
	}
	// 遍历长度最短的字符串，拼接返回值
	for _, currentChar := range targetStr {
		// 遍历全部字符
		isAllCheck := true
		for _, tmpStr := range A {
			if !isCharInString(currentChar, tmpStr) {
				isAllCheck = false
				break
			}
		}

		// 如果当前字符所有字符串都有
		if isAllCheck {
			result = append(result, string([]rune{currentChar}))
			for index := range A {
				A[index] = removeRuneFromStr(currentChar, A[index])
			}
		}
	}
	return result
}

//
// 判断字符串中是否有指定字符
//
func isCharInString(c rune, str string) bool {
	for _, char := range str {
		if char == c {
			return true
		}
	}
	return false
}

func removeRuneFromStr(c rune, str string) string {
	returnSlice := []rune{}
	hasDelete := false
	for _, tmpRune := range str {
		if tmpRune == c {
			if !hasDelete {
				hasDelete = true
				continue
			}
		}
		returnSlice = append(returnSlice, tmpRune)
	}
	return string(returnSlice)
}
