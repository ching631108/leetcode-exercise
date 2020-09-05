package main

import "fmt"

func main() {
	testData := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(testData, target)
	fmt.Println(result)
	return
}
