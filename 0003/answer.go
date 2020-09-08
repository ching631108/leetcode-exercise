package main

//
// 检查最长字符串长度
//
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	maxLonges := 0
	startIndex := 0
	for index := range s {
		tmpIndex := index + 1
		if checkRepeatStr(s[startIndex:tmpIndex]) {
			if (tmpIndex - startIndex) > maxLonges {
				maxLonges = tmpIndex - startIndex
			}
			continue
		}
		for i := startIndex + 1; i < tmpIndex; i++ {
			// 如果找到了
			startIndex = i
			if checkRepeatStr(s[i:tmpIndex]) {
				if (tmpIndex - i) > maxLonges {
					maxLonges = tmpIndex - i
				}
				break
			}
		}
	}
	return maxLonges
}

//
// 检查指定字符串中，是否存在重复字符
//
func checkRepeatStr(str string) bool {
	tmpMap := make(map[rune]bool)
	for _, char := range str {
		if tmpMap[char] {
			return false
		}
		tmpMap[char] = true
	}
	return true
}
