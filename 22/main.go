package main

import (
	"fmt"
	"strings"
)

// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var gener func(s string, l, r int)
	res := make([]string, 0)
	gener = func(s string, l, r int) {
		// 终止 len(s) = 2 * n
		// 左括弧 + 右括弧 是一对儿
		if len(s) == n*2 {
			res = append(res, strings.Clone(s))
			return
		}
		// 如果 l < n 选择 (
		if l < n {
			s += "("
			// 递归
			gener(s, l+1, r)
			// 还原
			s = s[:len(s)-1]
		}
		// 如果 r < l 选择 )
		if r < l {
			s += ")"
			gener(s, l, r+1)
			s = s[:len(s)-1]
		}
	}
	gener("", 0, 0)
	return res
}
