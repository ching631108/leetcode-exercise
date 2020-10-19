package main

//
// 基础方法
//
func backspaceCompareSimpleAnswer(S string, T string) bool {
	resultS := getResult(S)
	resultT := getResult(T)
	return resultS == resultT
}
func getResult(S string) string {
	arr := []rune{}
	for _, char := range S {
		if char != '#' {
			arr = append(arr, char)
		} else {
			if len(arr) > 0 {
				arr = arr[0 : len(arr)-1]
			}
		}
	}
	return string(arr)
}

//
// 双指针方法
//
func backspaceCompare(S string, T string) bool {
	sPointer := len(S) - 1
	tPointer := len(T) - 1

	// 两个字符串进行对比
	for sPointer >= 0 && tPointer >= 0 {
		sPointer = getNextPointer(S, sPointer)
		tPointer = getNextPointer(T, tPointer)
		if !(sPointer >= 0 && tPointer >= 0) {
			break
		}
		if S[sPointer] != T[tPointer] {
			return false
		}
		sPointer--
		tPointer--
	}
	// 两个字符串对比完成，如果还剩下字符串，则上下的字符串应该可以自己消化
	if sPointer >= 0 {
		sPointer = getNextPointer(S, sPointer)
	}
	if tPointer >= 0 {
		tPointer = getNextPointer(T, tPointer)
	}
	return sPointer < 0 && tPointer < 0
}
func getNextPointer(str string, pointer int) int {
	sharpCache := 0
	for {
		if pointer < 0 {
			return -1
		}
		if str[pointer] == '#' {
			sharpCache++
		} else {
			if sharpCache > 0 {
				sharpCache--
			} else {
				break
			}
		}
		pointer--
	}
	return pointer
}
