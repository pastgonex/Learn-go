package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	// 初始化一个数组
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// 创建一个切片， 对于arr[2:6] 的 view
	s := arr[2:6] // 2, 3, 4, 5
	fmt.Println("arr[2:6] = ", s)
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)

	// 修改slice的值相当于修改原数组对应切片位置的值
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)
}
