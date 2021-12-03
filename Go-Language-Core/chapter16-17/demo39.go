package main

import (
	"fmt"
)

func main() {
	num := 10
	sign := make(chan struct{}, num) // buffered channel

	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			// 这个struct{}{}占用的空间是0字节。
			// 确切地说，整个值在Go程序中永远都只会存在一份。
			// 虽然可以无数次使用，但是用的都是同一个值
			// 在channel队列中当做小成本占位符
			sign <- struct{}{} // struct{} 类型的（值）实例是struct{}{}

			/*
				我的理解是：在go func() 最后加一个 sign <- struct{}{}, 然后在主程中收
				这样可以保证每个go func() 都能执行完毕
			*/
		}()
	}

	// 方法1
	//time.Sleep(time.Millisecond * 500)

	// 方法2
	for j := 0; j < num; j++ {
		<-sign
	}
}
