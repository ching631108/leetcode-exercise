package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 处理预设字典
	replaceMap := map[string][]string{}
	for _, str := range wordList {
		starList := getStarList(str)
		for _, mapKey := range starList {
			replaceMap[mapKey] = append(replaceMap[mapKey], str)
		}
	}
	// fmt.Println(replaceMap)
	// 从初始节点开始，寻找通路

	// 记录当前寻找的层级
	currentLevel := 0
	currentLevelArr := []string{beginWord}
	historyStrArr := []string{}
	for !strInArr(endWord, currentLevelArr) {
		nextLevelArr := []string{}
		for _, tmpStr := range currentLevelArr {
			tmpStrList := getStarList(tmpStr)
			for _, nextStarStr := range tmpStrList {
				currentStrList := replaceMap[nextStarStr]
				for _, nextStr := range currentStrList {
					if !strInArr(nextStr, historyStrArr) {
						historyStrArr = append(historyStrArr, nextStr)
						nextLevelArr = append(nextLevelArr, nextStr)
					}
				}
			}
		}
		if len(nextLevelArr) == 0 {
			currentLevel = -1
			break
		}
		currentLevelArr = nextLevelArr
		currentLevel++
	}
	return currentLevel + 1
}

//
// 获取一个单词的 * 形式的全部数据
//
func getStarList(str string) []string {
	returnData := []string{}
	for strIndex := range str {
		tmpStr := []byte(str)
		tmpStr[strIndex] = '*'
		mapKey := string(tmpStr)
		returnData = append(returnData, mapKey)
	}
	return returnData
}

//
//	判断一个字符串是否属于数组
//
func strInArr(str string, arr []string) bool {
	for _, tmpStr := range arr {
		if str == tmpStr {
			return true
		}
	}
	return false
}
