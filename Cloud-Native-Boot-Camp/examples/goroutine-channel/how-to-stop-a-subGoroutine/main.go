package main

import "fmt"

// 如何停止一个子协程

func main() {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done channel is triggerred, exit child go routine")
			}
		}
	}()
	close(done)
}
