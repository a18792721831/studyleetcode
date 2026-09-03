package main

import (
	"fmt"
	"slices"
)

func main() {
	n := 4
	k := 2
	fmt.Println(combine(n, k))
}

// 输入：n = 4, k = 2
//输出：
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。

// 分析：n个数里面取k个数
// 这个题是选哪个
// n=4=> [1,2,3,4] 4个数里面选
// k=2 表示选2个
// [1,2] => 选 1 ,选 2
// [1,3] => 选 1 ,选 3
// [1,4] => 选 1 ,选 4
// [2,3] => 选 2 ,选 3
// [2,4] => 选 2 ,选 4
// [3,4] => 选 3 ,选 4
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var backtrack func(path []int, start int)
	backtrack = func(path []int, start int) {
		// 结束标志: len(path) == k
		if len(path) == k {
			res = append(res, slices.Clone(path))
			return
		}
		// 可行性剪枝
		// 如果是 n=5 k=3
		// 那么 当start = 4 就不需要继续了，没有意义
		for i := start; i <= n-k+len(path)+1; i++ {
			// 选 i
			path = append(path, i)
			// 递归
			backtrack(path, i+1)
			// 回退
			path = path[:len(path)-1]
		}
	}
	backtrack([]int{}, 1)
	return res
}
