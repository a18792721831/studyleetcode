package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[9,20],[15,7]]
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	// 队列
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)                // ★ 先存下当前层的规模，其实 q 里面是上一层的答案
		level := make([]int, 0, size) // level 用于收集上一层的结果
		for i := 0; i < size; i++ {   // 只出队 size 次——这一批就是完整一层
			n := q[0] // 获取 head
			q = q[1:] // 出队
			level = append(level, n.Val)
			if n.Left != nil { // 孩子入队——它们属于下一层
				q = append(q, n.Left) // 下一层迭代
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		// 上一层结果收集
		res = append(res, level)
	}
	return res
}
