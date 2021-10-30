// 寻找最长不含有重复字符的子串
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// 寻找最长不包含重复字符的子串
	var s string
	m := map[rune]int{}
	ans := 0
	fmt.Scan(&s)
	// str := []rune(s) // 转成rune即可支持中文
	// for i, j := 0, 0; i < len(str); i++ {
	// 	m[str[i]]++
	// 	for m[str[i]] > 1 {
	// 		m[str[j]]--
	// 		j++
	// 	}
	// 	ans = max(ans, i-j+1)
	// }
	j := 0
	str := []rune(s) // 转成rune即可支持中文
	for i, v := range str {
		m[v]++
		for m[v] > 1 {
			m[str[j]]--
			j++
		}
		ans = max(ans, i-j+1)
	}
	fmt.Println(ans)
}

