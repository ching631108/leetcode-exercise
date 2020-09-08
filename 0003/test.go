package main

import "fmt"

func main() {
	str := "abcabcabcbb"
	// isRepeat := checkRepeatStr(str[0:4])
	// fmt.Println(isRepeat)
	result := lengthOfLongestSubstring(str)
	fmt.Println(result)
	return
}
