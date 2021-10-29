package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	//result := ""
	//for ; n > 0; n /= 2 {
	//	lsb := n % 2
	//	result = strconv.Itoa(lsb) + result
	//}
	//return result
	result := ""
	for ; n > 0; n /= 2 {
		t := n % 2
		result += strconv.Itoa(t)
	}
	// reverse
	result = Reverse(result)
	return result
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func printFile(filename string) {
	// file, err := os.Open(filename)
	// if err != nil {
	// 	panic(err)
	// }
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

// 什么都不写， 就是一个死循环
// func forever() {
// 	for {
// 		fmt.Println("abc")
// 	}
// }

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1101
		convertToBin(72387885),
		convertToBin(1),
	)
	printFile("abc.txt")
	// forever()
}
