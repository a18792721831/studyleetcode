package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit(2, []int{2, 4, 1}))
	fmt.Println(maxProfit(1, []int{2, 4, 1}))
	fmt.Println(maxProfit(2, []int{3, 2, 6, 5, 0, 3}))
	fmt.Println(maxProfit(2, []int{1}))
}

// 输入：k = 2, prices = [2,4,1]
//输出：2
//解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。

//         3  2  6  5  0  3
// j = 0
// hold   -3 -2 -2  0  0  0
// cash    0  0  4  0  0  0
// j = 1
// hold   -m -2 -2 -1  4  4
// cash    0  0  4  4  4  4
// j = 2
// hold   -m -2 -2 -1  4  4
// cash    0  0  4  4  4  7
func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	if k == 0 {
		return 0
	}
	hold := make([][]int, k+1)
	for i := range hold {
		hold[i] = make([]int, n)
	}
	cash := make([][]int, k+1)
	for i := range cash {
		cash[i] = make([]int, n)
	}
	hold[0][0] = -prices[0]
	cash[0][0] = 0
	ans := -math.MaxInt
	for i := 1; i < n; i++ {
		hold[0][i] = max(hold[0][i-1], cash[0][i-1]-prices[i])
		cash[0][i] = 0
		ans = max(ans, cash[0][i])
	}
	for j := 1; j <= k; j++ {
		hold[j][0] = -math.MaxInt
		cash[j][0] = 0
		for i := 1; i < n; i++ {
			hold[j][i] = max(hold[j][i-1], cash[j][i-1]-prices[i])
			cash[j][i] = max(cash[j][i-1], hold[j-1][i-1]+prices[i])
			ans = max(ans, cash[j][i])
		}
	}
	return ans
}
