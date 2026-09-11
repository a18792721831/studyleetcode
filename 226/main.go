package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	// 分解式,先遍历，后处理
	if root == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
	// 遍历式，先处理，后遍历
	//if root == nil {
	//	return root
	//}
	//// 核心，交换左右子树
	//tmp := root.Left
	//root.Left = root.Right
	//root.Right = tmp
	//// 如果左子树不为叶子，继续
	//if root.Left != nil {
	//	invertTree(root.Left)
	//}
	//// 如果右子树不为叶子，继续
	//if root.Right != nil {
	//	invertTree(root.Right)
	//}
	//return root
}
