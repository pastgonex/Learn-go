package main

import (
	_ "Learn_go/Cloud-Native-Boot-Camp/examples/init/a"
	_ "Learn_go/Cloud-Native-Boot-Camp/examples/init/b"
	"fmt"
)

func init() {
	fmt.Println("main init")
}

func main() {
	fmt.Println("main")
}
