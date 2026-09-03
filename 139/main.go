package main

import (
	"fmt"
)

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
}

// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。

// 解题思路
// leetcode -> leet (选) -> leet(不选) | code(选)
// mem [idx] = 0 未知，1 可拆，2 不可拆
func wordBreak(s string, wordDict []string) bool {
	mem := make([]int8, len(s)+1)
	var backtrack func(idx int) bool
	backtrack = func(idx int) bool {
		// 结束条件 idx 到达尾端
		if idx == len(s) {
			return true
		}
		// 当前位置不可拆，直接剪枝
		if mem[idx] != 0 {
			return mem[idx] == 1
		}
		for i := range wordDict {
			// 存在子串
			if idx+len(wordDict[i]) <= len(s) && s[idx:idx+len(wordDict[i])] == wordDict[i] {
				if backtrack(idx + len(wordDict[i])) {
					mem[idx] = 1
					return true
				}
			}
		}
		mem[idx] = 2
		// 不存在任意子串
		return false
	}
	return backtrack(0)
}
