package main

func twoSum(nums []int, target int) []int {
	keyMap := make(map[int]int)
	for index, currentNumber := range nums {
		tmpKey := target - currentNumber
		if keyMap[tmpKey] == 0 {
			keyMap[currentNumber] = index + 1
			continue
		}
		return []int{keyMap[tmpKey] - 1, index}
	}
	return []int{0, 0}
}
