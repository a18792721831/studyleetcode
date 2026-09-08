package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}

//输入：[1,8,6,2,5,4,8,3,7]
//      0 1 2 3 4 5 6 7 8
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
// 双指针
// l=0,r=len-1
// l=1,r=7;1*8=8
// l=8,r=7;7*7=49 l+1
// l=1,r=3;1*7=7 r+1
// 直接简化 l 和 r ，谁大选谁
// l=6,r=7;6*6=36
// l=8,r=3;3*6=18
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	area := 0
	// 这个就不能排序了
	for ; l < r; {
		m := min(height[l], height[r])
		area = max(area, m*(r-l))
		// 面积都不大呢？
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return area
}
