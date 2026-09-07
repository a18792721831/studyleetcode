package main

import (
	"fmt"
)

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s, wordDict))
}

// 输入:s = "catsanddog", wordDict = ["cat","cats","and","sand","dog"]
// 输出:["cats and dog","cat sand dog"]
// dp解法
// dp[i] 表示单词 i 是否是s的子串
// 在进一步，dp[i][j] 表示在s中，从i到j是子串
//       c a t s a n d d o g
//     0 1 2 3 4 5 6 7 8 9 10
// can 1 0 0 1 1 0 0 1 0 0 1
//can2   1 0 0 1 1 0 0 1 0 0 1
// if
// can[i]=can[i-len(wordDict[x])

// 相当于构造了一个分割点表，将问题转换为 在 s 中插入空格
func wordBreak(s string, wordDict []string) []string {

	memo := make(map[int][]string) // memo[start] = 从 start 起能拼出的【所有句子】
	var dfs func(start int) []string
	dfs = func(start int) []string { // ← 有返回值！"回答问题"而不是"走路"
		if start == len(s) {
			return []string{""} // ← 哨兵：末尾产出"一个空句子"
		}
		if r, ok := memo[start]; ok { // 第二次到达：整棵子树一步不进
			return r
		}
		var out []string
		for _, w := range wordDict {
			end := start + len(w)
			if end <= len(s) && s[start:end] == w {
				for _, tail := range dfs(end) {
					if tail == "" {
						out = append(out, w) // 这个词就是整句
					} else {
						out = append(out, w+" "+tail) // 词 + 空格 + 后缀句子
					}
				}
			}
		}
		memo[start] = out // 缓存完整产出
		return out
	}
	return dfs(0)

	//can := make([]bool, len(s)+1)
	//can[0] = true
	//// 正序，can[i] 表示前缀 i 个 字符是否都在字典中
	//for i := 0; i < len(s)+1; i++ {
	//	for _, w := range wordDict {
	//		if i >= len(w) {
	//			can[i] = can[i] || can[i-len(w)] && s[i-len(w):i] == w
	//		}
	//	}
	//}
	//// can = 1 0 0 1 1 0 0 1 0 0 1
	//can2 := make([]bool, len(s)+1)
	//can2[len(s)] = true
	//// 倒序
	//for i := len(s); i >= 0; i-- {
	//	for _, w := range wordDict {
	//		if i+len(w) <= len(s) && !can2[i] {
	//			can2[i] = s[i:i+len(w)] == w && can2[i+len(w)]
	//		}
	//	}
	//}
	//// can2 = 1 0 0 1 1 0 0 1 0 0 1
	//res := make([]string, 0)
	//var backtrack func(path []string, idx int)
	//mem := make(map[int][]string, len(s)+1)
	//backtrack = func(path []string, idx int) {
	//	// 结束标志
	//	if idx == len(s) {
	//		res = append(res, strings.Join(path, " "))
	//		return
	//	}
	//	// 使用记忆
	//	if r, ok := mem[idx]; ok {
	//		// 选择
	//		for _, ss := range r {
	//			path = append(path, ss)
	//			// 递归
	//			backtrack(path, idx+len(ss))
	//			// 回退
	//			path = path[:len(path)-1]
	//		}
	//		return
	//	}
	//	for i := range wordDict {
	//		// 找到子串
	//		if idx+len(wordDict[i]) <= len(s) && s[idx:idx+len(wordDict[i])] == wordDict[i] && can2[idx+len(wordDict[i])] {
	//			mem[idx] = append(mem[idx], wordDict[i])
	//			// 选择
	//			path = append(path, wordDict[i])
	//			// 递归
	//			backtrack(path, idx+len(wordDict[i]))
	//			// 回退
	//			path = path[:len(path)-1]
	//		}
	//	}
	//	// 没找到任意子串
	//	return
	//}
	//backtrack([]string{}, 0)
	//return res

	//res := make([]string, 0)
	//var backtrack func(path []string, idx int)
	//mem := make(map[int][]string, len(s)+1)
	//backtrack = func(path []string, idx int) {
	//	// 结束标志
	//	if idx == len(s) {
	//		res = append(res, strings.Join(path, " "))
	//		return
	//	}
	//	// 使用记忆
	//	if r, ok := mem[idx]; ok {
	//		// 选择
	//		for _, ss := range r {
	//			path = append(path, ss)
	//			// 递归
	//			backtrack(path, idx+len(ss))
	//			// 回退
	//			path = path[:len(path)-1]
	//		}
	//		return
	//	}
	//	for i := range wordDict {
	//		// 找到子串
	//		if idx+len(wordDict[i]) <= len(s) && s[idx:idx+len(wordDict[i])] == wordDict[i] {
	//			mem[idx] = append(mem[idx], wordDict[i])
	//			// 选择
	//			path = append(path, wordDict[i])
	//			// 递归
	//			backtrack(path, idx+len(wordDict[i]))
	//			// 回退
	//			path = path[:len(path)-1]
	//		}
	//	}
	//	// 没找到任意子串
	//	return
	//}
	//backtrack([]string{}, 0)
	//return res
}
