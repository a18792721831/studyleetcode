package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	//fmt.Println(searchRange([]int{}, 0))
	fmt.Println(searchRange([]int{1}, 1))
}

//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
// 非递减，那就是等于和递增的数组
// 取中点 7 ，num[mid] < target , 意味着 [l,mid] 都小于target，[l,mid]都可以抛弃 l=mid+1
// 取中点 8 , 等于 target , 此时需要找第一个8，需要 l-- ，向左试探一格
// 第二个二分找最后一个8
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	l := 0
	r := len(nums) - 1
	// 找第一个target，本质上也是找点
	for ; l < r; {
		mid := (l + r) / 2
		if nums[mid] == target {
			// r指针移动到 mid
			// mid 有可能是第一个target
			// 至少可以说明 [mid+1,r]都是可以丢弃的
			r = mid
		} else if nums[mid] > target {
			// 取左边
			r = mid - 1
		} else {
			// 取右边
			l = mid + 1
		}
	}
	// 此时 l == r 一定是第一个 target
	// 如果target 不存在
	if nums[l] != target {
		return []int{-1, -1}
	}
	start := r
	r = len(nums) - 1
	// 还是找点，找右边的点
	for ; start <= r; {
		mid := (start + r) / 2
		if nums[mid] == target {
			// 找右边点，可以认为  [start,mid-1] 都是 target,都可以丢弃
			// mid 有可能是最后一个 target
			start = mid + 1
		} else if nums[mid] > target {
			// 取左边
			r = mid - 1
		} else {
			start = mid + 1
		}
	}
	// 此时 start 一定是最后一个 target
	return []int{l, start - 1}
}
