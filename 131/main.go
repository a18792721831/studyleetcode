package main

import (
	"fmt"
	"slices"
)

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
// a a b
// aa b
// aab 不对

func partition(s string) [][]string {
	res := make([][]string, 0)
	var gener func(path []string, idx int)
	gener = func(path []string, idx int) {
		// 终止条件 所有字符都用了
		if idx >= len(s) {
			res = append(res, slices.Clone(path))
			return
		}
		t_s := s[idx:]
		// 拆分子串的问题，按 1 位，2 位，3位拆分
		// 后续都任意拆分
		for i := 1; i <= len(t_s); i++ {
			t := t_s[:i]
			if !isNeed(t) {
				continue
			}
			// 选择
			path = append(path, t)
			// 递归
			gener(path, idx+i)
			// 回退
			path = path[:len(path)-1]
		}
	}
	gener([]string{}, 0)
	return res
}

// 判断是否是回文串
func isNeed(s string) bool {
	// 回文串的特点： 第一个字符等于最后一个字符
	if len(s) <= 1 {
		return true
	}
	// 奇数的话，最中间不管 aabaa
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
