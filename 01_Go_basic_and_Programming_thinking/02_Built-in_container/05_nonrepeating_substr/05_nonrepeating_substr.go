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
	m := map[byte]int{}
	ans := 0
	fmt.Scan(&s)
	for i, j := 0, 0; i < len(s); i++ {
		m[s[i]]++
		for m[s[i]] > 1 {
			m[s[j]]--
			j++
		}
		ans = max(ans, i-j+1)
	}
	fmt.Println(ans)
}
