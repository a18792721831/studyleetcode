package main

import "fmt"

func main() {
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
}

// 输入：s = "bbbab"
//输出：4
//解释：一个可能的最长回文子序列为 "bbbb" 。

// 选不选的问题，每一个元素都要判断选不选
// 区间 dp[i][j] 子串从 i .. j 最长的回文串长度
// 转义方程 s[i]==s[j] => dp[i][j]=dp[i+1][j-1]+2
//         s[i]!=s[j] => dp[i][j]=max(dp[i+1][j],dp[i][j-1])  不相等，所以不能给回文串的计数+1
//   b b b a b
//   0 1 2 3 4
// 0 1 2 3 3 4
// 1 0 1 2 2 3
// 2 0 0 1 1 2
// 3 0 0 0 1 1
// 4 0 0 0 0 1

//   c b b d
//   0 1 2 3
// 0 1 1 2 2
// 1 0 1 2 2
// 2 0 0 1 1
// 3 0 0 0 1

// 答案在右上角
// 行应该是从大到小
// 列应该是从小到大
func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			// 当前字符子串两端相等
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
