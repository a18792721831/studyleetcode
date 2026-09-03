package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

// 输入：[1,2,3,1]
//输出：4
//解释：偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
//     偷窃到的最高金额 = 1 + 3 = 4 。

// 动态规划类型题
// 解法：最后一个房间怎么处理？
// dp 数组记录的是当前房间可以偷的最大值
// i 偷，那么就是 dp[i]+dp[i-2]
// i 不偷，那么就是 dp[i-1]
func rob(nums []int) int {
	//if len(nums) == 1 {
	//	return nums[0]
	//}
	//if len(nums) <= 2 {
	//	if nums[0] > nums[1] {
	//		return nums[0]
	//	}
	//	return nums[1]
	//}
	//dp := make([]int, len(nums))
	//dp[0] = nums[0]
	//dp[1] = max(nums[1], dp[0])
	//for i := 2; i < len(nums); i++ {
	//	dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	//}
	//return dp[len(nums)-1]

	// 回溯版
	// 最小决策，偷不偷
	//maxV := 0
	//used := make([]bool, len(nums))
	//var backtrack func(sum, idx int)
	//backtrack = func(sum, idx int) {
	//	// 结束条件 idx == len(nums)
	//	if idx >= len(nums) {
	//		maxV = max(maxV, sum)
	//		return
	//	}
	//	// 合法性剪枝
	//	// 前一个已经偷了，当前不能再偷了
	//	if idx > 1 && used[idx-1] {
	//		backtrack(sum, idx+1)
	//		return
	//	}
	//	// 不偷
	//	backtrack(sum, idx+1)
	//	// 偷
	//	sum += nums[idx]
	//	used[idx] = true
	//	// 递归
	//	backtrack(sum, idx+1)
	//	// 还原
	//	sum -= nums[idx]
	//	used[idx] = false
	//}
	//backtrack(0, 0)
	//return maxV

	// 回溯记忆版
	// 记忆的是某个房间之前最大值
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	var backtrack func(idx int) int
	backtrack = func(idx int) int {
		// 结束条件
		if idx >= len(nums) {
			return 0
		}
		if memo[idx] >= 0 {
			return memo[idx]
		}
		memo[idx] = max(backtrack(idx+1), nums[idx]+backtrack(idx+2))
		return memo[idx]
	}
	return backtrack(0)
}
