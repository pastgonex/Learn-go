package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		ch1, ch2 := make(chan int), make(chan int)
		defer close(ch1)
		defer close(ch2)
		go func() {
			ch1 <- 1
		}()
		go func() {
			ch2 <- 2
		}()
		time.Sleep(time.Millisecond)
		select {
		case v := <-ch1:
			println("ch1:", v)
		case v := <-ch2:
			println("ch2:", v)
		default:
			println("nothing!")
		}
	*/

	messages := make(chan int, 10)
	done := make(chan bool)

	defer close(messages)
	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
