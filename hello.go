package main

import (
	"fmt"
	"strings"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	m := make(map[rune]int, 26)
	licensePlate = strings.ToLower(licensePlate)
	for _, v := range licensePlate {
		if 'a' <= v && v <= 'z' {
			m[v]++
		}
	}
	var ans string
	for _, s := range words {
		var cnt int
		for _, v := range s {
			if _, ok := m[v]; !ok {
				continue
			}
			cnt++
		}
		if cnt == len(s) {
			ans = s
			break
		}
	}
	return ans
}

func main() {
	fmt.Println("Hello world")
}
