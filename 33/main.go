package main

import "fmt"

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{1, 3}, 3))
	fmt.Println(search([]int{5, 1, 3}, 3))
}

// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4
// 二分查找：
// 如果是正常升序的数组 [0,1,2,4,5,6,7]
// 首先找中点 4 ,target < mid 取左边
// [0,1,2] 取中点 1, target < mid 取左边
// [0] == target 找到

// 这个题的麻烦点在于，进行了旋转，不是严格升序的
// 但是观察到虽然不是严格升序，但是依然是有序的，target 是分界点，target之前是严格升序，target之后也是严格升序
// [4,5,6,7,0,1,2] 取中点 7 ，同时取两段的左边界 4 和 0
// target < mid ,target < ll=4 ，取右段。
// 取右段的原因，左段的左边界是最小值，最小值大于target，那么不选
// [7,0,1,2,4,5,6] 如果是这种呢
// 将数组按照中点划分为两个数组，无法判定哪个是有序的 取哪边不一定
// [4,5,6,7,0,1,2]
// 按照中点切分，至少一半是有序的
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] != target {
			return -1
		}
		return 0
	}
	l := 0
	r := len(nums) - 1
	for ; l <= r; {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		// 左半边有序
		if nums[l] <= nums[mid] {
			// 取左半边
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				// 取右半边
				l = mid + 1
			}
		} else {
			// 右半边有序
			if nums[mid] < target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
