package main

type PartStr struct {
	start int
	end   int
}

func longestPalindrome(s string) string {
	// 特殊边界处理
	if len(s) <= 1 {
		return s
	}
	// 1. 寻找最小回文片段
	// 1.1. 寻找奇数回文片段
	oddSeed := getOddPalindrom(s)
	// 1.2. 寻找偶数回文片段
	evenSeed := getEvenPalindrom(s)
	allSeed := make([]PartStr, len(oddSeed)+len(evenSeed)+1)
	copy(allSeed[0:len(oddSeed)], oddSeed)
	copy(allSeed[len(oddSeed):len(oddSeed)+len(evenSeed)], evenSeed)
	allSeed = append(allSeed, PartStr{0, 1})
	// 2. 寻找最小回文片段对应的最大回文片段
	maxLength := 0
	returnStr := ""
	// 3. 寻找回文片段中最大的片段返回
	for _, currentSeed := range allSeed {
		tmpStr := getLongestPalindrome(s, currentSeed)
		if len(tmpStr) > maxLength {
			maxLength = len(tmpStr)
			returnStr = tmpStr
		}
	}
	return returnStr
}

// 获取奇数回文片段
func getOddPalindrom(s string) []PartStr {
	length := len(s)
	if length < 3 {
		return []PartStr{}
	}
	start := 0
	returnData := []PartStr{}
	for range s {
		if s[start] == s[start+2] {
			returnData = append(returnData, PartStr{
				start: start,
				end:   start + 3,
			})
		}
		start++
		if start >= length-2 {
			break
		}
	}
	return returnData
}

// 获取偶数回文片段
func getEvenPalindrom(s string) []PartStr {
	length := len(s)
	if length < 2 {
		return []PartStr{}
	}
	start := 0
	returnData := []PartStr{}
	for range s {
		if s[start] == s[start+1] {
			returnData = append(returnData, PartStr{
				start: start,
				end:   start + 2,
			})
		}
		start++
		if start >= length-1 {
			break
		}
	}
	return returnData
}

//
// 扩张指定字符串 寻找由该字符串开始的最长回文片段
//
func getLongestPalindrome(s string, seeder PartStr) string {
	start := seeder.start
	end := seeder.end
	length := len(s)
	for {
		if start < 1 {
			break
		}
		if end > length-1 {
			break
		}
		if s[start-1] != s[end] {
			break
		}
		start--
		end++
	}
	return s[start:end]
}
