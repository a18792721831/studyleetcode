package main

import "fmt"

func main() {
	fmt.Println(longestCommonSubsequence("ace", "abcde"))
	fmt.Println(longestCommonSubsequence("hofubmnylkra", "pqhgxgdofcvmr"))
}

// 输入：text1 = "abcde", text2 = "ace"
//输出：3
//解释：最长公共子序列是 "ace" ，它的长度为 3 。

// 解题思路：使用二维dp数组
// a => a b c d e
//      1 1 1 1 1
// c => 1 1 2 2 2
// e => 1 1 2 2 3
// dp[i][j]=max(dp[i][j-1],dp[i-1][j], dp[i-1][j-1]+1)

// dp[x][y] x < len(test1) y < len(test2)
// dp[x][y] 表示 text1 中，到 x 结尾，在 text2 中最长的公共序列
// abcde => x
// ace => y
// dp[5+1][3+1]
// dp[1][1] = 1
// dp[2][1] = 1
// dp[2][2] = 1
// dp[3][1] = 1
// dp[3][2] = 2
// dp[x][y] 的定义： 在text1中，以x结尾 与 text2 中，以y 结尾，最长的公共序列
// dp[3][3] = 2
// 推导： 双层循环，每次 x 步进 1 都需要比较看看text2 是否能接续是否存在
//

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
