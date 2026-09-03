package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}

// 输入：s = "25525511135"
// 输出：["255.255.11.135","255.255.111.35"]

func restoreIpAddresses(s string) []string {
	// 全局剪枝 [4,12]
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}
	res := make([]string, 0)
	var backtrack func(path []string, idx int)
	backtrack = func(path []string, idx int) {
		// 终止条件，所有字符都使用
		if idx >= len(s) && len(path) == 4 {
			res = append(res, strings.Join(path, "."))
			return
		}
		// 子串拆分问题，按照剩余数量拆分
		for i := 1; i <= len(s)-idx && i <= 3; i++ {
			str := s[idx : idx+i]
			// 前导 0 判断
			if len(str) > 1 && str[:1] == "0" {
				continue
			}
			// 拆分后的子串合法性剪枝
			ip, err := strconv.Atoi(str)
			// 非法字符
			if err != nil {
				continue
			}
			// 不合法
			if ip > 255 || ip < 0 {
				continue
			}
			// 可行性剪枝 本次拆完，剩下的 超出分配了
			if len(s)-idx-i > 3*(4-len(path)-1) {
				continue
			}
			// 选择
			path = append(path, str)
			// 递归
			backtrack(path, idx+i)
			// 回退
			path = path[:len(path)-1]
		}
	}
	backtrack([]string{}, 0)
	return res
}
