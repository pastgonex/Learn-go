package main

import "fmt"

func QuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	x := arr[(l+r)/2]
	i, j := l-1, r+1
	for i < j {
		for {
			i++
			if arr[i] >= x {
				break
			}
		}
		for {
			j--
			if arr[j] <= x {
				break
			}
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	QuickSort(arr, l, j)
	QuickSort(arr, j+1, r)
}

func main() {
	// arr := []int32{2, 3, 8, 7, 10}
	var n int
	arr := make([]int, n+10)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	QuickSort(arr, 0, n-1)
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d", arr[i])
	}
	fmt.Println()
}
