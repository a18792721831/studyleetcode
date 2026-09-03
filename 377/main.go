package main

import "fmt"

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}

// 输入：nums = [1,2,3], target = 4
//输出：7
//解释：
//所有可能的组合为：
//(1, 1, 1, 1)
//(1, 1, 2)
//(1, 2, 1)
//(1, 3)
//(2, 1, 1)
//(2, 2)
//(3, 1)
//请注意，顺序不同的序列被视作不同的组合。

// 顺序有关，所以是排列
// 外金额，内物品
func combinationSum4(nums []int, target int) int {
	if target == 0 {
		// 目标金额为0，空集合一种
		// 种子是1
		return 1
	}
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for _, c := range nums {
			if c <= i {
				dp[i] += dp[i-c]
			}
		}
	}
	return dp[target]
}
