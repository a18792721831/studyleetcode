package main

import "fmt"

func main() {

	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
}

//   2
// 0 0
// 1 -1
// 2 1
// 3 -1

// 输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1

// 完全dp
//    1  2 5
// 0  0  0 0
// 1  1 -1 -1
// 2  2  1 -1
// 3  3  2 -1
// 4  4  2 -1
// 5  5  3 1
// 6  6  3 2
// 7  7  4 2
// 8  8  4 3
// 9  9  5 3
// 10 10 5 2
// 11 11 6 3
// dp :=make([]int,amount+1)
// dp[0] 表示 凑 0 元，需要几个硬币，dp[0]=0
// dp[1] 表示 凑 1 元，需要几个硬币，dp[1]=1 // 有 1元 硬币
// dp[2] 表示 凑 2 元，需要几个硬币，dp[2]=1 // 因为有 2元 硬币
// dp[3] 表示 凑 3 元，需要几个硬币，dp[3]=dp[3-1]=1 dp[3-2]=1 取大还是小？不知道 => 取最小值+1 ？
// dp[4] 表示 凑 4 元，需要几个硬币，dp[4]=dp[4-1]=dp[3]=2 dp[4-2]=dp[2]=1
// dp[5] 表示 凑 5 元，需要几个硬币，dp[5]=1 // 因为有 5元 硬币
// dp[6] 表示 凑 6 元，需要几个硬币，dp[6]=dp[6-1]=1 dp[6-2]=dp[4]=2 dp[6-5]=dp[1]=1
// ....
// 结果在右下
func coinChange(coins []int, amount int) int {
	// 凑 0 元，0个硬币
	if amount == 0 {
		return 0
	}
	// 没有硬币，没法凑
	n := len(coins)
	if n == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			// 刚好有该面值的硬币
			if c <= i {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
