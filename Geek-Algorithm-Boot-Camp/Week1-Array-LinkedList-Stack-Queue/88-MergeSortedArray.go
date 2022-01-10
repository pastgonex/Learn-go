package main

// 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k], i = nums1[i], i-1
		} else {
			nums1[k], j = nums2[j], j-1
		}
		k--
	}
	// 就是在 nums1进行操作，没有意义
	// for i >= 0 {
	//     nums1[k], k, i = nums1[i], k-1, i-1
	// }
	for j >= 0 {
		nums1[k], k, j = nums2[j], k-1, j-1
	}
}

func mergeLydVersion(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for k >= 0 {
		if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
}
