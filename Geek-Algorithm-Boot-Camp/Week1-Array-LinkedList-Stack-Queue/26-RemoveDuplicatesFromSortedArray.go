package main

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}
