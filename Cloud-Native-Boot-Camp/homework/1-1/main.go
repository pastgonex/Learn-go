/*
作业要求：
	编写一个小程序
	给定一个字符串数组
	["I", "am", "stupid", "and", "weak"]
	用for循环遍历该数组并修改为
	["I", "am", "smart", "and", "strong"]
*/
package main

import "fmt"

func main() {
	str := [...]string{"I", "am", "stupid", "and", "weak"}
	for i := range str {
		if str[i] == "stupid" {
			str[i] = "smart"
		}
		if str[i] == "weak" {
			str[i] = "strong"
		}
	}
	fmt.Println(str)
}
