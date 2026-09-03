package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}

// 输入：candidates = [2,3,6,7], target = 7
//输出：[[2,2,3],[7]]
//解释：
//2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
//7 也是一个候选， 7 = 7 。
//仅有这两种组合。

// 解题分析
// 元素可以被重复使用
//
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	var backtrack func(path []int, sum int, start int)
	backtrack = func(path []int, sum int, start int) {
		// 结束条件 等于 target ，元素遍历完
		if sum == target {
			res = append(res, slices.Clone(path))
			return
		}
		for i := start; i < len(candidates); i++ {
			// 合法性剪枝 + 可行性剪枝，当前元素之和大于target，后面也一定大于
			if sum+candidates[i] > target {
				break
			}
			// 选择
			path = append(path, candidates[i])
			// 递归
			backtrack(path, sum+candidates[i], i)
			// 还原
			path = path[:len(path)-1]
		}
	}
	backtrack([]int{}, 0, 0)
	return res
}
