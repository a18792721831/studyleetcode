package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	// 中序遍历：左中右
	// 遍历，答案在中间
	// 只需要 bool 可以短路
	prev := -math.MaxInt
	var inorder func(r *TreeNode) bool
	inorder = func(r *TreeNode) bool {
		if r == nil {
			return true
		}
		if !inorder(r.Left) {
			return false
		}
		if r.Val <= prev {
			return false
		}
		// 中
		prev = r.Val
		// 右
		return inorder(r.Right)
	}
	return inorder(root)
}
