package main

import "fmt"

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}

// 输入：nums = [1,1,1,1,1], target = 3
//输出：5
//解释：一共有 5 种方法让最终目标和为 3 。
//-1 + 1 + 1 + 1 + 1 = 3
//+1 - 1 + 1 + 1 + 1 = 3
//+1 + 1 - 1 + 1 + 1 = 3
//+1 + 1 + 1 - 1 + 1 = 3
//+1 + 1 + 1 + 1 - 1 = 3

// + - 号，要不要，组合关系
// 物品在外，金额在内
// sum_p - sum_m = target => 1式
// sum_p + sum_m = sum(nums) => 2式
// 1式+2式 => 2 * sum_p = target + sum(nums)
// sum_p = (target + sum(nums)) / 2
// 这样就变成 要不要 ，凑 sum_p
// sum_p = (3 + 5)/2 = 4
// 变成 从nums中取 i 个数，和为 4 有多少种取法
// 1 1 1 1 1
// 0 1 1 1 1
// 1 0 1 1 1
// 1 1 0 1 1
// 1 1 1 0 1
// 1 1 1 1 0
// 每个数字必须选择要不要，所以是 0-1问题
// 倒序
func findTargetSumWays(nums []int, target int) int {
	// target == 0 是合法值
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	// 奇数和偶数处理
	if target > sum || target < -sum || (sum+target)%2 == 1 {
		return 0
	}
	n := (sum + target) / 2
	dp := make([]int, n+1)
	// 计数类， 种子为 1
	dp[0] = 1
	for _, c := range nums {
		for i := n; i >= c; i-- {
			dp[i] += dp[i-c]
		}
	}
	return dp[n]
}
