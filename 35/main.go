package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}

// 输入: nums = [1,3,5,6], target = 5
//输出: 2
//输入: nums = [1,3,5,6], target = 2
//输出: 1
//输入: nums = [1,3,5,6], target = 7
//输出: 4
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	l := 0
	r := len(nums) - 1
	// 二分找点
	for ; l <= r; {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
