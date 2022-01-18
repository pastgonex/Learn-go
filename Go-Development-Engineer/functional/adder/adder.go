package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	runAdder() // 累加器
}

func runAdder() {
	add := adder()
	fmt.Println("Start runAdder...")
	for i := 0; i < 10; i++ {
		if i > 2 {
			fmt.Printf("0 + 1 + ... + %d = %d\n", i, add(i))
		} else {
			fmt.Println(add(i))
		}
	}
	fmt.Println("runAdder done...")
}
