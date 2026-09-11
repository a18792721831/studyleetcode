package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	//fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	//fmt.Println(findKthLargest([]int{-1, 2, 0}, 2))
	//fmt.Println(findKthLargest([]int{7, 6, 5, 4, 3, 2, 1}, 5))
	//fmt.Println(findKthLargest([]int{8, 6, 5, 7}, 3))
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))

}

// 小顶堆
// 下沉是删除使用
// 上浮是插入使用
// 二叉树，如果是多叉树，那么就是 n*i+k
// 左孩子 ： 2*i+1
// 右孩子 ： 2*2+2
// 父亲 ： (i-1)/2
func findKthLargest(nums []int, k int) int {
	// 快速排序
	//输入: [3,2,1,5,6,4], k = 2
	//输出: 5
	// 假设输入是有序的，并且从小到大，那么倒数第2大的值，就是从尾部往回数2个
	// pivot=len-k=4
	// 先选取末尾元素4
	// 从左往右找大于4的值
	// 3,2,1,5 > 4 ,5和4交换位置 [3 2 1 4 6 5]
	// 此时 pivot 已经归位 ,在 3 < 4 继续右边，因为左边都小于 pivot,last=5
	// 6 > 5 ,交换位置 [3 2 1 4 5 6] last=6
	// 二路 lomuto 可能会超时，三路 lomuto 是正解之一。
	// 小顶堆是最优解
	lomuto := func(nums []int, l, r int) int {
		pivot := nums[r] // ① pivot 钉死：全程没人动 nums[r]
		i := l           // ② i 是分界线：i 左边全是 < pivot 的
		for j := l; j < r; j++ {
			if nums[j] < pivot { // ③ j 扫描：遇到小的就换到 i 区，分界线右扩
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[r] = nums[r], nums[i] // ④ 归位：pivot 插进分界线
		return i                            // pivot 的最终下标
	}
	target := len(nums) - k
	l, r := 0, len(nums)-1
	for {
		p := lomuto(nums, l, r)
		if p == target {
			return nums[p]
		} else if p < target {
			// p 小于 target ,说明目标在右边
			l = p + 1
		} else {
			// p 大于 target,说明目标在左边
			r = p
			l = 0
		}
		idx := l + rand.Intn(r-l+1)
		nums[idx], nums[r] = nums[r], nums[idx] // 随机选一个换到末尾，再照常 Lomuto

	}
	//if len(nums) == 0 {
	//	return 0
	//}
	//// tok k 大，使用小顶堆，堆顶是最小值。
	//myheap := make([]int, 0)
	//// 建立堆
	//for i := 0; i < len(nums); i++ {
	//	// 先要把堆装满
	//	// 和堆顶比
	//	if len(myheap) >= k && nums[i] < myheap[0] {
	//		// 比最小值还小，没有进入堆的条件
	//		continue
	//	}
	//	// 堆不满，直接加末尾，然后上浮
	//	if len(myheap) < k {
	//		// 装堆
	//		// 元素加入数组末尾
	//		myheap = append(myheap, nums[i])
	//		index := len(myheap) - 1
	//		fIndex := (index - 1) / 2
	//		// 到达根节点，或者父亲比我小
	//		for ; myheap[fIndex] > myheap[index]; fIndex = (index - 1) / 2 {
	//			// 父亲比我大，那么需要交换位置
	//			t := myheap[index]
	//			myheap[index] = myheap[fIndex]
	//			myheap[fIndex] = t
	//			index = fIndex
	//		}
	//	} else {
	//		// 堆已经满了
	//		// 更新堆顶，然后下沉
	//		myheap[0] = nums[i]
	//		// 要找孩子节点
	//		index := 0
	//		// 左孩子
	//		lc := index*2 + 1
	//		// 右孩子
	//		rc := index*2 + 2
	//		// 我要比所有孩子小
	//		for (lc < len(myheap) && myheap[index] > myheap[lc]) || (rc < len(myheap) && myheap[index] > myheap[rc]) {
	//			// 我要和最小的孩子交换位置
	//			if rc >= len(myheap) {
	//				// 没有右孩子,直接和左孩子交换
	//				t := myheap[lc]
	//				myheap[lc] = myheap[index]
	//				myheap[index] = t
	//				index = lc
	//			} else if myheap[lc] < myheap[rc] {
	//				// 有右孩子，且左孩子最小
	//				// 最大的孩子是左节点
	//				t := myheap[lc]
	//				myheap[lc] = myheap[index]
	//				myheap[index] = t
	//				index = lc
	//			} else {
	//				// 右孩子大
	//				t := myheap[rc]
	//				myheap[rc] = myheap[index]
	//				myheap[index] = t
	//				index = rc
	//			}
	//			lc = index*2 + 1
	//			rc = index*2 + 2
	//		}
	//	}
	//}
	//// 返回堆顶
	//return myheap[0]
}
