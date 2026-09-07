package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

// 输入：prices = [7,1,5,3,6,4]
//输出：7
//解释：在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5 - 1 = 4。
//随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6 - 3 = 3。
//最大总利润为 4 + 3 = 7 。

//       7  1  5  3  6  4
// hold -7 -1 -1  1  1  3
// cash  0  0  4  4  7  7
//            这里能获利，就卖出
// 但是这个我感觉是不对的，因为数字顺序变了，不对的。
// 隐含一个条件，只要cash大于0，就完成一次交易
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	hold := make([]int, n)
	cash := make([]int, n)
	hold[0] = -prices[0]
	cash[0] = 0
	for i := 1; i < n; i++ {
		hold[i] = max(hold[i-1], cash[i-1]-prices[i])
		cash[i] = max(cash[i-1], hold[i]+prices[i])
	}
	return cash[n-1]
}
