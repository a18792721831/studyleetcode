package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
	fmt.Println(twoSum([]int{1, 2, 3, 4, 4, 9, 56, 90}, 8))
	fmt.Println(twoSum([]int{-5, -3, 0, 2, 4, 6, 8}, 5))
}

//输入：numbers = [2,7,11,15], target = 9
//输出：[1,2]
//解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。
// 观察：非递减，那就是等于+递增
// 观察：要求常数级别的算法
// 双指针，一左一右。
// 左指针表示小值，右指针表示大值。
// 左指针从0开始，有指针从中间开始
// 和小于target，右指针移动(更大)，和大于target，左指针移动(更小)
// 注意：下标从1开始
func twoSum(numbers []int, target int) []int {
	if len(numbers) == 2 {
		// 题目说明一定有结果，当数组长度等于2，一定是答案
		return []int{1, 2}
	}
	l := 0
	r := len(numbers) - 1
	// 条件是不能越界
	for ; l < r; {
		if numbers[l]+numbers[r] == target {
			// 下标从 1开始
			break
		} else if numbers[l]+numbers[r] > target {
			// 右指针不能再动了，要找一个更小的值
			r--
		} else {
			l++
		}
	}
	return []int{l + 1, r + 1}
}
