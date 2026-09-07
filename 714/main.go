package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}

// 输入：prices = [1, 3, 2, 8, 4, 9], fee = 2
//输出：8
//解释：能够达到的最大利润:
//在此处买入 prices[0] = 1
//在此处卖出 prices[3] = 8
//在此处买入 prices[4] = 4
//在此处卖出 prices[5] = 9
//总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8
//       1  3   2  8  4  9
// hold -1 -1  -1 -1  1  1
// cash  0  0   0  5  5  8
// hold[i] = max(hold[i-1], -prices[i]
// cash[i] = max(cash[i-1], hold[i]+prices[i]-2]
func maxProfit(prices []int, fee int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	hold := make([]int, n)
	cash := make([]int, n)
	hold[0] = -prices[0]
	cash[0] = 0
	for i := 1; i < n; i++ {
		hold[i] = max(hold[i-1], cash[i-1]-prices[i])
		cash[i] = max(cash[i-1], hold[i]+prices[i]-fee)
	}
	return cash[n-1]
}
