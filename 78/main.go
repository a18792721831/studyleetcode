package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

// 输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 分析：[] -> 不选，不选，不选
// [1] -> 选，不选，不选
// [1,2] 等价于 [2,1] 无序的，需要同层剪枝
// 首先，排序，利用 start 处理同层剪枝
func subsets(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	var backtrack func(path []int, start int)
	backtrack = func(path []int, start int) {
		// 结束条件，start 走到结束
		if start >= len(nums) {
			res = append(res, slices.Clone(path))
			return
		}
		// 开始循环
		// 不选
		backtrack(path, start+1)
		// 选
		path = append(path, nums[start])
		// 递归
		backtrack(path, start+1)
	}
	backtrack([]int{}, 0)
	return res
}
