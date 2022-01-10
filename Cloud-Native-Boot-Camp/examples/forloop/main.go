package main

import "fmt"

func main() {
	//for i := 0; i < 10; i++ {
	//	println(i)
	//}
	fullString := "hello world"
	fmt.Println(fullString)
	for i, c := range fullString {
		fmt.Println(i, string(c))
	}
}
