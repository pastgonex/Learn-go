package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer func is called")
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()
	panic("a panic is triggered")
}
