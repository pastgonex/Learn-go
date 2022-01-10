package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Microsecond * 35) // 为协程设定超时时间
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	select {
	// check normal channel
	case <-ch:
		fmt.Println("received from ch")
	case <-timer.C:
		fmt.Println("timeout waiting from channel ch")
	}
}
