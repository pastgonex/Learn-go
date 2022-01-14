package main

import "fmt"

func main() {
	// panic这种级别的错误, 是需要recover的
	defer func() {
		fmt.Println("defer func is called")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("a panic is triggered!!")
}
