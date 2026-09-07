package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

// 输入: prices = [1,2,3,0,2]
// 输出: 3
// 解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

//hold[i]     = max(hold[i-1], free[i-1] - prices[i])   ← 买入只能从"非冷冻的空仓"来
//justSold[i] = hold[i-1] + prices[i]                      ← 今天卖出（明天冷冻）
//free[i]     = max(free[i-1], justSold[i-1])              ← 可买状态：一直空着 or 冷冻解禁
// 手推
//        1  2  3  0 2
// hold  -1 -1 -1  1 1
// sold  -m  1  2 -1 3
// free   0  0  1  2 2
// sold 吸收了冷静期，卖出的时候，会用到sold

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	hold := make([]int, n)
	cash := make([]int, n)
	sold := make([]int, n)
	hold[0] = -prices[0]
	cash[0] = 0
	sold[0] = -math.MaxInt
	for i := 1; i < n; i++ {
		hold[i] = max(hold[i-1], cash[i-1]-prices[i])
		sold[i] = hold[i-1] + prices[i]
		cash[i] = max(cash[i-1], sold[i-1])
	}
	return max(cash[n-1], sold[n-1])
}
