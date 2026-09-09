package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
	fmt.Println(findMin([]int{3, 1, 2}))
	fmt.Println(findMin([]int{5, 1, 2, 3, 4}))
}

//输入：nums = [3,4,5,1,2]
//输出：1
//解释：原数组为 [1,2,3,4,5] ，旋转 3 次得到输入数组。

// 取中点 5
// 3 < 5 左半边有序 左半边最小 3
// 右半边无序，右半边最大值 2 < 3 最小值在右半边
// 1 2
// 中点 = 1

//输入：nums = [4,5,6,7,0,1,2]
//输出：0
//解释：原数组为 [0,1,2,4,5,6,7] ，旋转 4 次得到输入数组。

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	for ; l < r; {
		mid := (l + r) / 2
		// 最小值在右边，左边全部丢弃
		if nums[mid] >= nums[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
