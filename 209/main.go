package main

import "fmt"

func main() {
	//fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	//fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	//fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	//fmt.Println(minSubArrayLen(15, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}))
	fmt.Println(minSubArrayLen(5, []int{2, 3, 1, 1, 1, 1, 1}))
}

// 输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 滑动窗口题
// l,r
// window
// [2 3 1 2 4 3]
// l =0,r=l+1
// window [2 3] sumWindow=5 < target r++
// window [2 3 1] sumWindow=6 < target r++
// window [2 3 1 2] sumWindow=8 > target , sumWindow-l=6 < target，所以 不能移动l,r++ ,count = 4
// window [2 3 1 2 4] sumWindow=12 > target ,sumWindow-l=10 > target ,继续移动l,sumWindow-l=7, 继续移动,sumWindow-l=6<target,不满足，l=2 count=3,r++
// window [1 2 4 3] sumWindow=10 > target ,sumWindow -l=9 > target,继续移动l,sumWindow-l=7,继续移动，sumWindow-l=3<target,不满足，l=4,count=2,r++
func minSubArrayLen(target int, nums []int) int {
	window := 0
	l := 0
	r := 0
	count := 0
	ans := len(nums)
	for ; r < len(nums); {
		// 先加入 window
		window += nums[r]
		count++
		r++
		if window >= target {
			ans = min(ans, count)
		}
		for ; window >= target && window-nums[l] >= target; {
			// 移出window
			window -= nums[l]
			count--
			ans = min(ans, count)
			l++
		}
	}
	// 全部加起来也不够
	if window < target {
		return 0
	}
	return ans
}
