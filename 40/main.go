package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(candidates, 8))
}

// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]

// 核心决策，选哪个
// 排序
// used
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	var backtrack func(path []int, start int, sum int)
	sort.Ints(candidates)
	backtrack = func(path []int, start int, sum int) {
		// 结束条件
		// sum = target
		// 结果集里面没有
		if sum == target {
			res = append(res, slices.Clone(path))
			return
		}
		for i := start; i < len(candidates); i++ {
			// 同层去重
			if i > start && candidates[start] == candidates[i] {
				continue
			}
			// 当前元素加上后大于 target 剪枝
			if sum+candidates[i] > target {
				break
			}
			// 当前元素加上后小于 target ,选择
			path = append(path, candidates[i])
			// 递归 每次向后移动一个
			backtrack(path, i+1, sum+candidates[i])
			// 恢复
			path = path[:len(path)-1]
		}
	}
	backtrack([]int{}, 0, 0)
	return res
}
