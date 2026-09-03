package main

import "fmt"

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

// 输入：word1 = "horse", word2 = "ros"
//输出：3
//解释：
//horse -> rorse (将 'h' 替换为 'r')
//rorse -> rose (删除 'r')
//rose -> ros (删除 'e')

// 是不是可以这么理解：ros 和 horse 倒序比较，删除 horse 的 e 和 将 horse 的 e 替换为 s 然后在 删除前一个 s 相比，直接删除 e 操作步数最少
//       "" h o r s e
// "" =>  0 1 2 3 4 5 ：从 "" 空字符串到 上面的 "",h,ho,hor,hors,horse 各自需要几步，几个字母就需要插入几次
// r  =>  1 1 2 2 3 4 : 从  r 到 "",h,ho,hor,hors,horse => 1,1,2,2,3,4
// o  =>  2 2 1 2 3 4 : 从 ro 到 "",h,ho,hor,hors,horse => 2,2,1,2,3,4
// s  =>  3 3 2 2 2 3 : 从 ros 到  "",h,ho,hor,hors,horse => 3,3,2,2,2,3
//        : 竖着的，表示从 "",r,ro,ros 需要几步到 "" 。反着来也行
// 将
// dp 定义： dp[i][j] 将字符串1与字符串2 ，从开始到下标 i j 替换相同，需要操作的最少步数
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 1; i < n+1; i++ {
		dp[0][i] = i
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				// 末尾字符相等，不变
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[m][n]
}
