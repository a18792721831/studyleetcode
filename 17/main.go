package main

import (
	"fmt"
	"strings"
)

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

// 输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

// 解题分析：
// 整体还是选哪个，但是有个映射: 2 - [a,b,c] 3 - [d,e,f] ....
// 输入 2 3
// [ad] - 选 a 选 d
// [ae] - 选 a 选 e
// 输入是一个字符串
// 输出是一个字符串数组
func letterCombinations(digits string) []string {
	res := make([]string, 0)
	var backtrack func(path []string, start int)
	backtrack = func(path []string, start int) {
		// 结束条件： start == len(dig)
		if start >= len(digits) {
			res = append(res, strings.Join(path, ""))
			return
		}
		// 选择，这里有个有趣的点，每个按钮只会对应3个字母
		for i := 0; i < 4; i++ {
			// 7 和 9 例外
			if i == 3 && digits[start:start+1] != "7" && digits[start:start+1] != "9" {
				break
			}
			// 选择
			path = append(path, sw(digits[start:start+1], i))
			// 递归
			backtrack(path, start+1)
			// 回退
			path = path[:len(path)-1]
		}
	}
	backtrack([]string{}, 0)
	return res
}

func sw(s string, idx int) string {
	var str string
	switch s {
	case "2":
		str = "abc"
	case "3":
		str = "def"
	case "4":
		str = "ghi"
	case "5":
		str = "jkl"
	case "6":
		str = "mno"
	case "7":
		str = "pqrs"
	case "8":
		str = "tuv"
	case "9":
		str = "wxyz"
	}
	return str[idx : idx+1]
}
