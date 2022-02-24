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

func removeDuplicatesEverNi(nums []int) int {
	cnt, n := 1, len(nums)
	for i, j := 0, 1; i < n; {
		j = i + 1
		for j < n && nums[i] == nums[j] {
			j++
		}
		if j >= n {
			break
		}
		nums[cnt] = nums[j]
		i, cnt = j, cnt+1
	}
	return cnt
}
