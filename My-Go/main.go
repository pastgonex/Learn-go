package main

import "fmt"

type point struct {
	x, y, len int
}

func initPoint(p *point) {
	*p = point{1, 2, 3}
}

func main() {
	p := &point{}
	initPoint(p)
	fmt.Println(p.x)
}
