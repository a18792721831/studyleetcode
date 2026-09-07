package main

import (
	"fmt"
)

func main() {
	//fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2}))
}

//输入：[7,1,5,3,6,4]
//输出：5
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。

// 解法：
// dp[i] 当前价格卖出，可以获得的最大利润
// 种子 = 0
// 倒序会更好
// 手推
//     7 1 5 3 6 4
//   0 1 2 3 4 5 6
// 1 7 0 0 0 0 0 0
// 2 1 0 0 4 2 5 3
// 3 5 0 0 0 0 1 0
// 4 3 0 0 0 0 3 1
// 5 6 0 0 0 0 0 0
// 6 4 0 0 0 0 0 0
// 答案在中间
// 换个角度，算差值
//      7  1  5  3  6  4
//   0  1  2  3  4  5  6
// 1 7  0 -6 -2 -4 -1 -3
// 2 1 -6  0 -4 -2 -5 -3
// 3 5 -2  4  0  2 -1  1
// 4 3 -4  2 -2  0 -3 -1
// 5 6 -1  5  1  3  0  2
// 6 4 -3  3 -1  1 -2  0

// 相当于需要两个dp
// 第一个hold 表示持有成本，找最大
// 第二个需要借助第一个，表示利润最大
// 答案在 cash 最后
//
//       7  1  5  3  6  4
// hold -7 -1 -1 -1 -1 -1
// cash  0  0  4  4  5  5
// hold[0] = -max
// hold[i] = max(hold[i-1],-prices[i])
// cash[0] = 0
// cash[i] = max(cash[i-1], hold[i] +prices[i])

//        1  2
// hold  -1 -1
// cash   0  1
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	buy, sale := prices[0], 0
	for i := 0; i < n; i++ {
		sale = max(sale, prices[i]-buy)
		buy = min(buy, prices[i])
	}
	return sale

	//n := len(prices)
	//if n == 0 {
	//	return 0
	//}
	//hold := make([]int, n)
	//cash := make([]int, n)
	//hold[0] = -prices[0]
	//cash[0] = 0
	//for i := 1; i < n; i++ {
	//	hold[i] = max(hold[i-1], -prices[i])
	//	cash[i] = max(cash[i-1], hold[i]+prices[i])
	//}
	//return cash[n-1]
}
