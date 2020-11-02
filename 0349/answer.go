package main

func intersection(nums1 []int, nums2 []int) []int {
	uniqueNums1 := uniqueSlice(nums1)
	uniqueNums2 := uniqueSlice(nums2)
	result := []int{}
	for _, n1 := range uniqueNums1 {
		isRepeat := false
		for _, n2 := range uniqueNums2 {
			if n1 == n2 {
				isRepeat = true
				break
			}
		}
		if isRepeat {
			result = append(result, n1)
		}
	}
	return result
}

func uniqueSlice(nums []int) []int {
	numsLen := len(nums)
	result := []int{}
	for j := 0; j < numsLen; j++ {
		hasRepeat := false
		for k := j + 1; k < numsLen; k++ {
			if nums[k] == nums[j] {
				hasRepeat = true
				break
			}
		}
		if !hasRepeat {
			result = append(result, nums[j])
		}
	}
	return result
}
