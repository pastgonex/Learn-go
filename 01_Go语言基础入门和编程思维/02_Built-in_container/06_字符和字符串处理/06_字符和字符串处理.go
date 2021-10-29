package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我爱你"
	fmt.Println(len(s))
	fmt.Printf("%x\n", []byte(s))
	for _, b := range []byte(s) { // utf8编码 中文3字节
		fmt.Printf("%x ", b) // utf-8编码
	}
	fmt.Println()

	// rune 将 string进行utf-8解码，转换成 unicode
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d, %x) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	utf8.DecodeRune(bytes)

}