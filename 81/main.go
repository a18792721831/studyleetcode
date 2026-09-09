package main

import "fmt"

func main() {
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	//fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	//fmt.Println(search([]int{1, 0, 1, 1, 1}, 0))
	//fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{3, 5, 1}, 1))
}

// 输入：nums = [2,5,6,0,0,1,2], target = 0
//输出：true
//输入：nums = [2,5,6,0,0,1,2], target = 3
//输出：false
// 1, 0, 1, 1, 1 ,target=0

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return true
		}
		return false
	}
	l := 0
	r := len(nums) - 1
	for ; l < r; {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		// 右边有序
		if nums[l] == nums[mid] {
			l++
		} else if nums[r] == nums[mid] {
			r--
		} else if nums[r] > nums[mid] {
			// 右边最小值大于target,target在左边
			if nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 左边有序
			// 左边最大值小于target,target在右边
			if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
