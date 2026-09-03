package main

import "fmt"

func main() {
	//fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	fmt.Println(findLength([]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}))
}

// 输入：nums1 = [1,2,3,2,1], nums2 = [3,2,1,4,7]
//输出：3
//解释：长度最长的公共子数组是 [3,2,1] 。

// 有点像上一个比较的子串的问题
//   0 3 2 1 4 7
// 0 0 0 0 0 0 0
// 1 0 0 0 1 0 0
// 2 0 0 1 0 0 0
// 3 0 1 0 0 0 0
// 2 0 0 2 0 0 0
// 1 0 0 0 3 0 0

//   0 0 1 1 1 1
// 0 0 0 0 0 0 0
// 1 0 0 1 1 1 1
// 0 0 1 0 0 0 0
// 1 0 0 2 1 1 1
// 0 0 1 0 0 0 0
// 1 0 0 2 1 1 1
// 这里有个问题，不连续的怎么处理？

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	ans := 0
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}
