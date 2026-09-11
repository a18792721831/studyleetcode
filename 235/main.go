package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//if root == nil || root == q || root == p {
	//	return root
	//}
	// 因为是 BST ，所以可以根据值进行判断
	// 核心逻辑： p,q < ro 往左找，p,q > ro 往右找， p<ro<q 返回ro
	//if p.Val < root.Val && q.Val < root.Val {
	//	return lowestCommonAncestor(root.Left, p, q)
	//} else if (p.Val < root.Val && q.Val > root.Val) || (q.Val < root.Val && p.Val > root.Val) {
	//	return root
	//} else {
	//	return lowestCommonAncestor(root.Right, p, q)
	//}
	for root != nil {
		if root == p || root == q {
			return root
		}
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return root
}
