package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// 如果改为后序呢？
	// 中序：左中右
	// 后序：左右中
	//stack := []*TreeNode{}
	//cur := root
	//for len(stack) > 0 || cur != nil {
	//	for cur != nil { // ① 一路向右，全部压栈
	//		// 要改成后序，必须先入栈中节点，然后入栈右节点，在入栈左节点
	//		// root->left1->left11->left111
	//		stack = append(stack, cur)
	//		cur = cur.Left
	//	}
	//	// 注意点，这里是从切片尾端拿一个
	//	// 最左节点 left111
	//	cur = stack[len(stack)-1] // ② 弹出栈顶 = 当前中序节点
	//	stack = stack[:len(stack)-1]
	//	// 左节点取出，但是不处理，先看右节点
	//	for cur.Right != nil { // ① 一路向右，全部压栈
	//		// 要改成后序，必须先入栈中节点，然后入栈右节点，在入栈左节点
	//		// root -> left1 -> left11 -> right111
	//		// 如果右节点比左节点长呢？
	//		// root->left1->left11->right111->left1111->right1111
	//		stack = append(stack, cur.Right)
	//		cur = cur.Right
	//	}
	//	k--
	//	if k == 0 {
	//		return cur.Val // ← 裸 break：栈在你手里，提前停就是 return
	//	}
	//
	//}
	//return 0

	// 如果从中序改为前序呢？
	// 中序：左中右
	// 前序：中左右
	// 前序迭代版：弹即处理 + 逆序压子
	//stack := []*TreeNode{root}
	//for len(stack) > 0 {
	//	cur := stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	k--
	//	if k == 0 {
	//		return cur.Val          // 前序位置：弹出即处理，无延迟
	//	}
	//	if cur.Right != nil {
	//		stack = append(stack, cur.Right)   // 右先压
	//	}
	//	if cur.Left != nil {
	//		stack = append(stack, cur.Left)    // 左后压 → 左先弹（LIFO）
	//	}
	//}

	stack := []*TreeNode{}
	cur := root
	for len(stack) > 0 || cur != nil {
		for cur != nil { // ① 一路向左，全部压栈
			// 先入栈，在左节点入栈
			// root -> left1 -> left11 -> left111
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 注意点，这里是从切片尾端拿一个
		// 最左节点 left111
		cur = stack[len(stack)-1] // ② 弹出栈顶 = 当前中序节点
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return cur.Val // ← 裸 break：栈在你手里，提前停就是 return
		}
		// 最左子树的右节点入栈
		// root->left1->left11->right111
		// 如果right111有子树，在从左节点开始
		cur = cur.Right // ③ 转向右子树
	}
	return 0
	//q := make([]int, 0)
	//var inorder func(ro *TreeNode)
	//inorder = func(ro *TreeNode) {
	//	if ro == nil || k <= 0 {
	//		return
	//	}
	//	inorder(ro.Left)
	//	k--
	//	q = append(q, ro.Val)
	//	inorder(ro.Right)
	//}
	//inorder(root)
	//return q[len(q)-1]

	//// BST，中序 k
	//var ans int
	//var inorder func(ro *TreeNode)
	//inorder = func(ro *TreeNode) {
	//	if ro == nil || k <= 0 {
	//		return
	//	}
	//	inorder(ro.Left)
	//	k--
	//	if k == 0 {
	//		ans = ro.Val
	//		return
	//	}
	//	inorder(ro.Right)
	//}
	//return ans
}
