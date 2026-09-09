package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("S"))
	fmt.Println(lengthOfLongestSubstring("mq"))
	fmt.Println(lengthOfLongestSubstring("1R1T7"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

// 输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
// 0 <= s.length <= 10^5
// 手推
//     a b c a b c b b
//     1 2 3 3 3 3 2 1
//     p w w k e w
//     1 2 1 2 3 3
// left
// dp[i]=i-left+1
// s[i]=s[left]

// 输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
// 滑动窗口题
// window = [128]int
// {a},{ab},{abc},{abca},{bca},{bcab},{cab},{cabc},{abc},{abcb},{cb},{cbb},{b}
// 窗口移出的条件，当 窗口中存在重复字符的时候，此时顺序移出字符，直到不存在重复字符
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	l := 0
	r := 0
	ans := 0
	win := make([]int, 128)
	for ; r < len(s); {
		// 进窗口
		win[s[r]]++
		// 有重复字符
		for ; win[s[r]] > 1 && l < r; {
			// 移出窗口
			win[s[l]]--
			l++
		}
		r++
		ans = max(ans, r-l)
	}
	return ans
	//left := 0
	//last := make([]int, 128)
	//last[s[0]] = 1
	//ans := 1
	//for i := 1; i < len(s); i++ {
	//	if last[s[i]] > left {
	//		left = last[s[i]]
	//	}
	//	ans = max(ans, i-left+1)
	//	last[s[i]] = i + 1
	//}
	//return ans
}
