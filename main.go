package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
Scan会不停地调用Read方法， 直到Read遇到EOF
要想要intGen类型当做File来使用， 就可以给他实现一个Read方法，并在某种逻辑中返回EOF
这个intGen.Read方法， 会在Scan的时候调用， 并传入一个buffer（具体的切片索引）
*/

func genFib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

// Read就是读到EOF停止， 一直读意味着 next会不断生成
// p是一个总的buffer，Scan的时候会往里面塞东西
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF // 需要给他一个io.EOF, 否则会一直读（因为读不到EOF）
	}
	//return next, nil
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

// 只要实现了Read方法(Reader接口), 那么就可以当成是 file 了
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() { // Scan应该会不停地调用Read方法
		fmt.Println(scanner.Text())
	}
}

func main() {
	f := genFib()
	printFileContents(f)
}
