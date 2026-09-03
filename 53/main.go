package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

//输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
//输出：6
//解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

// 以i结尾的最大值
// -2 和 1 ，1 重新开始是最大
// dp = [-2 1 -2 4 3 5 6 1 5]
func maxSubArray(nums []int) int {
	//if len(nums) == 0 {
	//	return 0
	//}
	//if len(nums) == 1 {
	//	return nums[0]
	//}
	//dp := make([]int, len(nums))
	//dp[0] = nums[0]
	//ans := dp[0]
	//for i := 1; i < len(nums); i++ {
	//	dp[i] = max(dp[i-1]+nums[i], nums[i])
	//	ans = max(ans, dp[i])
	//}
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	curr := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		curr = max(curr+nums[i], nums[i])
		ans = max(curr, ans)
	}
	return ans
}
