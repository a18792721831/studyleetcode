package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(solveNQueens(4))
}

// 输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。

// 皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
// 这里其实隐含一个条件，就是对角线不能放棋子
// 最小决策，当前格子是否能放棋子 ,继续拆分，应该是这一层的棋子，放哪个位置
// 需要记住哪一行或那一列不能再放
// r+c:      r-c:
//0 1 2 3      0 -1 -2 -3
//1 2 3 4      1  0 -1 -2
//2 3 4 5      2  1  0 -1
//3 4 5 6      3  2  1  0
// 对斜线划分了编号 r+c : 0,6
// r-c +n-1 : 0-6
func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	res := make([][]string, 0)
	// 直接搞一个二维的byte数组
	pan := make([][]byte, n)
	// 初始化棋盘，全部 .
	for i := 0; i < n; i++ {
		pan[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			pan[i][j] = '.'
		}
	}
	// 这里需要三个使用的标记
	// 列标记
	used := make([]bool, n)
	// 左上到右下的标记
	dig1 := make([]bool, 2*n-1)
	// 左下到右上的标记
	dig2 := make([]bool, 2*n-1)
	var queue func(row int)
	queue = func(row int) {
		// 结束条件
		if row == n {
			str := make([]string, n)
			for i := 0; i < n; i++ {
				str[i] = string(pan[i])
			}
			res = append(res, slices.Clone(str))
			return
		}
		for i := 0; i < n; i++ {
			if used[i] || dig1[row-i+n-1] || dig2[row+i] {
				continue
			}
			// 选择
			used[i] = true
			dig1[row-i+n-1] = true
			dig2[row+i] = true
			pan[row][i] = 'Q'
			// 递归
			queue(row + 1)
			// 还原
			used[i] = false
			dig1[row-i+n-1] = false
			dig2[row+i] = false
			pan[row][i] = '.'
		}
	}
	queue(0)
	return res
}
