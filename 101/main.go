package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//输入：root = [1,2,2,3,4,4,3]
//输出：true
func isSymmetric(root *TreeNode) bool {
	// 轴对称，左等于右
	// 自顶向下
	if root == nil {
		return true
	}
	var compare func(left, right *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		if left == right {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return compare(left.Left, right.Right) && compare(left.Right, right.Left)
	}
	return compare(root.Left, root.Right)
}
