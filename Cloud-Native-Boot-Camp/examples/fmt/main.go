package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	fmt.Printf("%s\n", "Never give up")
	_ = fmt.Errorf("%s\n", "an error!")
	str := fmt.Sprintf("%s+%s", "QWQ", "nbq")
	fmt.Println(str)
}
