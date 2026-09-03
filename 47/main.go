package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}

// 输入：nums = [1,1,2]
// 输出：
// [[1,1,2],
//  [1,2,1],
//  [2,1,1]]
// 这个是选哪个，不是要不要的问题
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	// 排序
	sort.Ints(nums)
	isUsed := make([]bool, len(nums))
	var backtrack func(path []int, isUsed []bool, idx int)
	backtrack = func(path []int, isUsed []bool, idx int) {
		// 终止条件
		if len(path) == len(nums) {
			res = append(res, slices.Clone(path))
			return
		}
		for i := 0; i < len(nums); i++ {
			// 使用过的元素不要
			if isUsed[i] {
				continue
			}
			// 横向判断
			// 1 2 2 里面，如果 前面 2 不要，后面的 2 也不应该要
			// 因为 1 2 2 里面，如果前面的 2 要，后面 2 不要，已经包含了 1 2
			// 在 1 2 2 里面，前面的 2 不要，后面的 2 要，就重复了
			if i > 0 && nums[i] == nums[i-1] && !isUsed[i-1] {
				continue
			}
			// 选择
			isUsed[i] = true
			path = append(path, nums[i])
			// 递归
			backtrack(path, isUsed, idx+1)
			// 回退
			isUsed[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack([]int{}, isUsed, 0)
	return res
}
