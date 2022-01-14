package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 148560
	}()

	timer := time.NewTimer(time.Microsecond * 40) // 为协程设定超时时间
	select {
	// check normal channel
	case v := <-ch:
		fmt.Printf("value %d received from ch\n", v)
	case <-timer.C:
		fmt.Println("timeout waiting from channel ch")
	}
}
