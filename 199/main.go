package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			t := q[0]
			q = q[1:]
			// 收集结果
			if i == size-1 {
				res = append(res, t.Val)
			}
			// 孩子节点处理
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
	}
	return res
}
