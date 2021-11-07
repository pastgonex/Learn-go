package main

import "fmt"

func main() {
	s := make([]int, 2)
	fmt.Println(s)
	s = append(s, 1)
	s = append(s, 2)
	fmt.Println(s)
	s = append(s, 3)
	fmt.Println(s)

	m := make(map[int]int, 16)
	fmt.Printf("%T, %v", m, m)
}
