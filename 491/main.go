package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println(findSubsequences(nums))
}

// 输入：nums = [4,6,7,7]
//输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]

// 解题思路：
// 选不选的模式
// 约束：1. 至少2个元素
// 2. 递增
// 3. 同层剪枝， [4 6 7_1] 和 [4 6 7_2] 是相同的
func findSubsequences(nums []int) [][]int {
	// 全局剪枝
	if len(nums) < 2 {
		return [][]int{}
	}
	res := make([][]int, 0)
	var path []int
	var backtrack func(start int)
	backtrack = func(start int) {
		// 结束条件 start >= len
		if len(path) >= 2 {
			res = append(res, slices.Clone(path))
		}
		mem := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			// 数字已经用过了
			if mem[nums[i]] {
				continue
			}
			// 合法性剪枝 + 可行性剪枝
			if len(path) > 0 && path[len(path)-1] > nums[i] {
				continue
			}
			// 选择
			path = append(path, nums[i])
			mem[nums[i]] = true
			// 递归
			backtrack(i + 1)
			// 还原
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
