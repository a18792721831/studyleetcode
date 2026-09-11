package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//输入：root = [1,2,3,4,5]
//输出：3
//解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
func diameterOfBinaryTree(root *TreeNode) int {
	// 直径是两个节点之间最长的路劲
	// 换句话说就是左子树最长+右子树最长
	// 答案可能在中间，不一定经过root根节点
	var depth func(ro *TreeNode) int
	ans := 0
	depth = func(ro *TreeNode) int {
		if ro == nil {
			return 0
		}
		left := depth(ro.Left)
		right := depth(ro.Right)
		ans = max(ans, left+right)
		return max(left, right) + 1
	}
	depth(root)
	return ans
}
