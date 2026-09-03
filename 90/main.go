package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}

// 输入：nums = [1,2,2]
// 输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]

// 最小决策，取不取的问题,取 0 个，1个，2个
// 排序,同层剪枝
// 取不取在一次循环里都做
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	used := make([]bool, len(nums))
	path := make([]int, 0)
	var backtrack func(idx int)
	backtrack = func(idx int) {
		// 结束条件，idx 到了结尾
		// idx 标识了取几个
		if idx == len(nums) {
			res = append(res, slices.Clone(path))
			return
		}
		// 不要 分支
		backtrack(idx + 1)

		// 要分支
		// 1 2 里面，需要同层剪枝
		if idx > 0 && nums[idx] == nums[idx-1] && !used[idx-1] {
			return
		}
		// 选择
		path = append(path, nums[idx])
		used[idx] = true
		// 递归
		backtrack(idx + 1)
		// 恢复
		used[idx] = false
		path = path[:len(path)-1]
	}
	backtrack(0)
	return res
}
