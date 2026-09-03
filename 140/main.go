package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s, wordDict))
}

// 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
// 输出:["cats and dog","cat sand dog"]
func wordBreak(s string, wordDict []string) []string {
	res := make([]string, 0)
	var backtrack func(path []string, idx int)
	mem := make(map[int][]string, len(s)+1)
	backtrack = func(path []string, idx int) {
		// 结束标志
		if idx == len(s) {
			res = append(res, strings.Join(path, " "))
			return
		}
		// 使用记忆
		if r, ok := mem[idx]; ok {
			// 选择
			for _, ss := range r {
				path = append(path, ss)
				// 递归
				backtrack(path, idx+len(ss))
				// 回退
				path = path[:len(path)-1]
			}
			return
		}
		for i := range wordDict {
			// 找到子串
			if idx+len(wordDict[i]) <= len(s) && s[idx:idx+len(wordDict[i])] == wordDict[i] {
				mem[idx] = append(mem[idx], wordDict[i])
				// 选择
				path = append(path, wordDict[i])
				// 递归
				backtrack(path, idx+len(wordDict[i]))
				// 回退
				path = path[:len(path)-1]
			}
		}
		// 没找到任意子串
		return
	}
	backtrack([]string{}, 0)
	return res
}
