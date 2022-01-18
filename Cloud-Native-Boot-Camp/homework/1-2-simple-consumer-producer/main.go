package main

import "time"

/*
基于 Channel 编写一个简单的单线程生产者消费者模型
队列：
队列长度10， 队列元素类型为int
生产者：
每1秒往队列中放入一个类型为int的元素， 队列满时生产者可以阻塞
消费者：
每1秒从队列中获取一个元素并打印， 队列为空时消费者阻塞
*/

func main() {
	ch := make(chan int, 10)
	timer := time.NewTimer(time.Second * 1)
	for {
		select {
		case <-timer.C:
			timer.Reset(time.Second * 1)
			ch <- 1
		case v := <-ch:
			println(v)
		}
	}
}
