package main

// 移动零
func moveZeroes(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			n++
		}
	}
	for n < len(nums) {
		nums[n], n = 0, n+1
	}
}
