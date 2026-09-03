package main

import (
	"fmt"
)

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{3, 3, 3, 4, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}

// 输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。

// 如果是回溯的题，就是选不选，不选，当前元素进入第二个数组，选，当前元素进入第一个数组，要求，分完了，两个数组之和相等
// 因为最终结果是bool，使用快速短路返回
// 而且最后只判断和，所以不需要数组，只需要 sum1和sum2分别代表和

// 问题，采用回溯写，会超时

// dp解法
// 如果dp是bool
// 1 5 11 5
// 我看到了提示：dp[i][j]
// 还是用二维的数组，dp[i][j] 表示 前 i 个元素之和是否等于 j
//    i 1 5 11 5
// j  t t t t  t
// 1  f t t t  t
// 2  f f f f  f
// 3  f f f f  f
// 4  f f f f  f
// 5  f f t t  t
// 6  f f t t  t
// 7  f f f f  f
// 8  f f f f  f
// 9  f f f f  f
// 10 f f f f  t
// 11 f f f t  t
// 答案在最后。

func canPartition(nums []int) bool {
	m := len(nums)
	if m == 0 {
		return true
	}
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	n := sum / 2
	dpn := make([]bool, n+1)
	dpn[0] = true
	for i := 1; i < m+1; i++ {
		for j := n; j >= 0; j-- {
			if j >= nums[i-1] {
				// 选择分支
				dpn[j] = dpn[j] || dpn[j-nums[i-1]]
			}
		}
	}
	return dpn[n]
	//
	//dp := make([][]bool, m+1)
	//for i := range dp {
	//	dp[i] = make([]bool, n+1)
	//	dp[i][0] = true
	//}
	//for i := 1; i < m+1; i++ {
	//	for j := 1; j < n+1; j++ {
	//		if j >= nums[i-1] {
	//			// 选择分支
	//			dp[i][j] = dp[i-1][j-nums[i-1]]
	//		}
	//		// 不选分支
	//		dp[i][j] = dp[i][j] || dp[i-1][j]
	//
	//	}
	//}
	//return dp[m][n]

	//// 正好可以复习一遍
	//var backtrack func(idx, sum1, sum2 int) bool
	//backtrack = func(idx, sum1, sum2 int) bool {
	//	if idx >= len(nums) {
	//		if sum1 == sum2 {
	//			return true
	//		}
	//		return false
	//	}
	//	// 选择 要
	//	sum1 += nums[idx]
	//	// 递归
	//	if backtrack(idx+1, sum1, sum2) {
	//		// 快速短路返回
	//		return true
	//	}
	//	// 还原
	//	sum1 -= nums[idx]
	//	// 选择 不要
	//	sum2 += nums[idx]
	//	// 递归
	//	if backtrack(idx+1, sum1, sum2) {
	//		return true
	//	}
	//	// 还原
	//	sum2 -= nums[idx]
	//	return false
	//}
	//return backtrack(0, 0, 0)
}
