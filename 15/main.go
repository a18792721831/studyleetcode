package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}

//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//解释：
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
//不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
//注意，输出的顺序和三元组的顺序并不重要。

// 解题思路：
// 先取一个数，然后从剩余的里面取两个数，这两个数的和等于第一个数
// 顺序无关，排序加下标处理
// -1 0 1 2 -1 -4
// -4 -1 -1 0 1 2
// 先取 -4 剩余取两个数，达不到 -4
// 取 -1, 剩余两个数 -1,2
// 第二个-1 ，同层剪枝
// 取 0 剩余2个数，-1,1
// 取 1 剩余两个数，-1,0 -> 这里是重复的，怎么处理？
// 取 2 剩余两个数，-1,-1 -> 这里也是重复的，怎么处理？
// 用回溯的话，可以使用used表处理，但是规模限制，肯定不能使用回溯
// 先选一个数，剩余数从剩余数字中取
// 双指针
// -4 -1 -1 0 1 2
// -4:
// l=-1,r=2;l+r=1<4;l++
// l=-1,同层剪枝;l++
// l=0,r=2;l+r=2<4;l++
// l=1,r=2;l+r=4<4;l++;l==r ->结束
// -1:
// l=-4,r=2;l+r=-2<1;l++
// l=-1 = target ,l++
// l=-1 同层剪枝;l++ -> 这里被误伤了
// l=0,r=2;l+r=2>1;r--
// l=0,r=1;l+r=1 -> 答案;l++,r-- ;l>= r -> 结束
// -1: 同层剪枝,idx++
// 0:
// l=-4,r=2;l+r=-2<0;l++
// l=-1,r=2;l+r=1>0;r--
// l=-1,r=1;l+r=0;答案;l++,r--
// l=-1,同层剪枝,l++
// l==idx;l++;l>=r -> 结束
// 1:
// l=-4,r=2;l+r=-2<-1;l++
// l=-1,r=2;l+r=1 >-1;r--
// r==idx,r--
// l=-1,r=0;l+r=-1==target;答案;l++,r--;l>=r -> 结束
// 2:
// l=-4,r=1;l+r=-3 < -2;l++
// l=-1,r=1;l+r=0>-2;r--
// l=-1,r=0;l+r=-1>-2;r--
// l=-1,r=-1;l+r=-2==target 答案;l++,r--;l>=r -> 结束
// 答案：
// -1,0,1;0,-1,1;1,-1,0;2,-1,-1
// 为何答案有重复？
// 如果限制 l>=idx?

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		// 外层的同层剪枝
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 全正数不能最终和为0
		if nums[i] > 0 {
			break
		}
		for ; l < r; {
			if nums[l]+nums[r] == -nums[i] {
				res = append(res, []int{nums[l], nums[r], nums[i]})
				for ; l < r && nums[l] == nums[l+1]; {
					l += 1
				}
				for ; l < r && nums[r] == nums[r-1]; {
					r -= 1
				}
				l += 1
				r -= 1
				continue
			}
			if nums[l]+nums[r] > -nums[i] {
				r -= 1
				continue
			}
			if nums[l]+nums[r] < -nums[i] {
				l += 1
				continue
			}
		}
	}
	return res
}
