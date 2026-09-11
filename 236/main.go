package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
	//if root == nil {
	//	return nil
	//}
	//var cla func(ro *TreeNode) *TreeNode
	//cla = func(ro *TreeNode) *TreeNode {
	//	if ro == nil {
	//		return nil
	//	}
	//	left := cla(ro.Left)
	//	right := cla(ro.Right)
	//	if ro == p {
	//		return p
	//	}
	//	if ro == q {
	//		return q
	//	}
	//	if left != nil && right != nil {
	//		return ro
	//	}
	//	if left != nil {
	//		return left
	//	}
	//	if right != nil {
	//		return right
	//	}
	//	return nil
	//}
	//return cla(root)
}
