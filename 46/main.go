package main

import "fmt"

func main() {
	nums1 := []int{5, 4, 6, 2}
	fmt.Println(permute(nums1))
}

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var backtrack func(lastNums []int)
	backtrack = func(lastNums []int) {
		if len(lastNums) == len(nums) {
			n := make([]int, len(lastNums))
			copy(n, lastNums)
			res = append(res, n)
			return
		}
		// 1,2,3
		for idx := range nums {
			if isHave(lastNums, nums[idx]) {
				continue
			}
			// 1
			lastNums = append(lastNums, nums[idx])
			// 2,3
			backtrack(lastNums)
			lastNums = lastNums[:len(lastNums)-1]
		}

	}
	backtrack([]int{})
	return res
}

func isHave(nums []int, i int) bool {
	for _, v := range nums {
		if v == i {
			return true
		}
	}
	return false
}
