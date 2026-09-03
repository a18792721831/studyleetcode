package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isAdditiveNumber("112358"))
	fmt.Println(isAdditiveNumber("199100199"))
	fmt.Println(isAdditiveNumber("111"))

}

// 输入："112358"
//输出：true
//解释：累加序列为: 1, 1, 2, 3, 5, 8 。1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8

// 解题思路：这是一个包含字符串分割，和分割后处理的复杂问题
// 首先，这是一个短路问题
// 输入的字符串，进行分割，需要处理前导0
// 先是字符串，如何拆分有效的序列
func isAdditiveNumber(num string) bool {
	// 全局剪枝
	if len(num) < 3 {
		return false
	}
	// start 表示位置，结束标志
	var backtrack func(path []int, start int) bool
	backtrack = func(path []int, start int) bool {
		// 结束条件
		if start >= len(num) && len(path) > 2 {
			// 刚好拆分完了
			return true
		}
		for i := 1; i+start <= len(num); i++ {
			// 获取拆分的子串
			str := num[start : start+i]
			// 0 是允许的，但是 01 不允许
			if len(str) > 1 && str[:1] == "0" {
				// 前导 0,整枝剪掉 01 和 011 都是前导0
				break
			}
			// 子串转数字
			is, ok := strconv.Atoi(str)
			if ok != nil {
				// 数字转换出错
				return false
			}
			// 前两位直接选
			if len(path) < 2 {
				// 选择
				path = append(path, is)
				// 从下一位开始
				// 需要利用短路
				if backtrack(path, start+i) {
					return true
				}
				// 还原
				path = path[:len(path)-1]
				continue
			}
			// 从第三位开始要进行合法性剪枝
			// 不合法整体剪枝
			if path[len(path)-1]+path[len(path)-2] == is {
				// 选择
				path = append(path, is)
				// 递归
				return backtrack(path, start+i)
			}
			// 可行性剪枝
			if is > path[len(path)-1]+path[len(path)-2] {
				break
			}
		}
		return false
	}
	return backtrack([]int{}, 0)
}
