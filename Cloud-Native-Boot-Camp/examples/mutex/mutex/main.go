package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Mutex: 互斥锁
- Mutex为互斥锁，Lock()加锁， Unlock()解锁
- 在一个goroutine获得Mutex后， 其他goroutine只能等到这个goroutine释放该mutex
- 使用Lock()加锁后，不能再继续对其加锁， 直到利用Unlock()解锁后才能再加锁
- 在Lock()之前使用Unlock()会导致panic
- 已经锁定的Mutex并不与特定的goroutine相关联，这样可以利用一个goroutine对其加锁
  再利用其他的goroutine对其解锁
- 在同一个goroutine中的Mutex解锁之前再次进行进行加锁，会导致死锁
- 适用于读写不确定，并且只有一个读或者写的场景
*/

func main() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("the lock is locked")

	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not lock:", i)
			mutex.Lock()
			fmt.Println("Locked:", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock:", i)
			mutex.Unlock()
			c <- i
		}(i, channels[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range channels {
		<-c
	}

}
