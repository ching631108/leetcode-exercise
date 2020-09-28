package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 6}
	a2 := []int{2, 5, 6, 7, 8}
	result := findMedianSortedArrays(a1, a2)
	fmt.Println(result)

	return
}
