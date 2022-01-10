package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	/*
		for i := 0; i < 10; i++ {

			// 如果不加go的话， 这段程序是在同一个CPU处理的
			// 加了go之后， 是告诉程序，起新的CPU使用， 所以顺序是没有保证的

			go fmt.Println("loopFunc:", i)
		}

		// 这种make的时候，不加缓冲通道长度的，默认是0，因此你塞了一个进去，就阻塞了，直到有人取出来
		ch := make(chan int)
		go func() {
			fmt.Println("hello from goroutine")
			ch <- 999
		}()
		i := <-ch // 从Channel中取数据并赋值
		fmt.Println(i)
	*/
	ch := make(chan int, 10)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			go func() {
				rand.Seed(time.Now().UnixNano())
				n := rand.Intn(10)
				ch <- n
			}()
		}
		time.Sleep(time.Millisecond) // 在这里sleep是因为要在这个协程结束之前，等待一会
		//close(ch)
	}()
	fmt.Println("hello from main")
	for v := range ch {
		fmt.Println("receiving:", v)
	}

}
