package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	loopFunc()
	time.Sleep(time.Second)
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		// 如果这里不使用闭包的话（将这块逻辑圈起来）
		go func(i int) {
			lock.Lock()
			defer lock.Unlock() // defer 就类似于 finally
			fmt.Println("loopFunc:", i)
		}(i)
	}
}
