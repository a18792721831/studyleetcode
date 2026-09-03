package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
}

// 输入：nums = [10,9,2,5,3,7,101,18]
//输出：4
//解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

// 这里需要记一下当前最大元素
// dp 数组，当前元素结尾的最长递增子序列的长度
// 9 < 10 不递增，且 max(dp[i-1], 1) = 1 max = 9
// 2 < 9 不递增，且 max(dp[i-1], 1) = 1 max = 2
// 5 > 2 递增，取 max(dp[i-1]+1,1) = 2 max = 5 max=nums[i]=5
// 3 < 5 不递增，取 max(dp[i-1],1) = 2 max = 5
// 7 > 5 递增，取 max(dp[i-1]+1,1) = 3 max = 7
// 101 > 7 递增，取 max(dp[i-1]+1,1) = 4 max = 101
// 18 < 101 不递增，取 max(dp[i-1],1) = 4

// 10,9,2,5,3,7,101,18
// 1,1,1,2,2,2,3,4,4
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
