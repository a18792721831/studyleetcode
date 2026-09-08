package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("a", "aa"))
	fmt.Println(minWindow("a", "b"))
	fmt.Println(minWindow("aa", "aa"))
}

// 输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
//解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。

// ADOBECODEBANC
// ABC
// need['A']=1,need['B']=1,need['C']=1
// window
// count=0
// l=0,r=l+1 {AD},window['A']=1,count=1
// l=0,r=2 {ADO}
// {ADOB} window['B']=1,count=2
// {ADOBEC} window['C']=1,count=3 -> l=1,r=5; ans=ADOBEC ,window['A']=0
// {DOBECODEBA} window['A']=1,count=3,window['B']=2,window['C']=1 -> l++
// {ECODEBA} window['B']=1,count=3 - >
// 只要count = len(need) 就可以移动 l
// {ODEBA} window['C']=0,count=2 ,r++
// {ODEBANC} window['C']=1,count=3,l++
// {ANC} window['B'] = 0 count=2,r++ r> len(s)结束
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	need := make(map[byte]int)
	for idx := range t {
		need[t[idx]]++
	}
	window := make(map[byte]int)
	l := 0
	r := 0
	count := 0
	ans := ""
	for ; r < len(s); {
		if need[s[r]] > 0 {
			window[s[r]]++
			if window[s[r]] == need[s[r]] {
				count++
			}
		}
		r++
		for ; count == len(need); {
			if need[s[l]] > 0 {
				window[s[l]]--
				if window[s[l]] < need[s[l]] {
					count--
					if len(s[l:r]) < len(ans) || len(ans) == 0 {
						ans = s[l:r]
					}
				}
			}
			l++
		}
	}
	return ans
}
