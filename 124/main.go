package main

import "math"

func main() {
	t1 := &TreeNode{
		Val: 1,
	}
	t2 := &TreeNode{
		Val: 2,
	}
	t3 := &TreeNode{
		Val: 3,
	}
	t1.Left = t2
	t1.Right = t3
	maxPathSum(t1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 不一定经过根节点，需要在中间收集答案
func maxPathSum(root *TreeNode) int {
	// 路径和= 左子树和+当前节点+右子树和
	// 因为找最大，所以需要不断取max
	if root == nil {
		return 0
	}
	ans := -math.MaxInt
	var gian func(ro *TreeNode) int
	gian = func(ro *TreeNode) int {
		if ro == nil {
			return 0
		}
		left := max(gian(ro.Left), 0)
		right := max(gian(ro.Right), 0)
		ans = max(ans, ro.Val+left+right)
		return ro.Val + max(left, right)
	}
	gian(root)
	return ans
}
