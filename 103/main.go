package main

import (
	"slices"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	q := []*TreeNode{root}
	levelIndex := 1
	for len(q) > 0 {
		size := len(q)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			t := q[0]
			q = q[1:]
			// 收集值
			level = append(level, t.Val)
			// 处理孩子
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		// 奇数 从右往左
		if levelIndex%2 == 0 {
			slices.Reverse(level)
		}
		res = append(res, level)
		levelIndex++
	}
	return res
}
