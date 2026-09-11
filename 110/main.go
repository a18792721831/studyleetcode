package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 输入：root = [3,9,20,null,null,15,7]
//输出：true
// 平衡二叉树，左右子树高度差小于等于1
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// depth 只返回深度
	// 是否平衡由 ans 判断
	var depth func(ro *TreeNode) int
	depth = func(ro *TreeNode) int {
		if ro == nil {
			return 0
		}
		left := depth(ro.Left)
		right := depth(ro.Right)
		if left == -1 || right == -1 || !(left-right <= 1 && left-right >= -1) {
			return -1
		}
		return max(left, right) + 1
	}
	return depth(root) != -1
}
