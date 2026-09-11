package main

import "slices"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			// 收集结果
			t := q[0]
			q = q[1:]
			level = append(level, t.Val)
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		res = append(res, level)
	}
	slices.Reverse(res)
	return res
}
