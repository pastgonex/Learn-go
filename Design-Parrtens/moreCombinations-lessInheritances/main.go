// 多用组合 少用继承
// 虽然go中没有继承， 但这里说明了组合的好处
package main

import "fmt"

type interface1 interface {
	a() string
}

type interface2 interface {
	b() int
}

// 组合让接口中的代码 没有必要实现很多份
type struct1 struct {
	struct2 // 结构体的组合， 直接使用struct2中的field 和 method
	x, y    int
}

type struct2 struct {
	xxx, yyy int
}

func (s struct1) a() string {
	return "called Method in interface1 "
}
func (s struct2) b() int {
	return 999
}

func main() {
	s1 := struct1{}
	fmt.Println(s1.b())
	fmt.Println(s1.xxx)
}
