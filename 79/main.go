package main

import "fmt"

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word))
	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "SEE"
	fmt.Println(exist(board, word))
	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word = "ABCB"
	fmt.Println(exist(board, word))

}

// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCCED"
// 输出：true
// 解题思路：选哪个，可以认为有四个方向，上下左右，同时不能重复使用
func exist(board [][]byte, word string) bool {
	if len(board) < 1 {
		return false
	}
	if len(word) < 1 {
		return true
	}
	var dfs func(x, y int, start int) bool
	m := len(board)
	n := len(board[0])
	// 因为不能重复使用，所以需要记录是否已经在所选列表内
	used := make([][]bool, len(board))
	// 初始化
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}
	dfs = func(x, y int, start int) bool {
		// 可行性剪枝
		if x >= m || x < 0 || y >= n || y < 0 {
			// 撞墙了
			return false
		}
		if used[x][y] || board[x][y] != word[start] {
			return false
		}
		if len(word)-1 == start {
			return true
		}
		used[x][y] = true
		found := dfs(x+1, y, start+1) || dfs(x-1, y, start+1) || dfs(x, y+1, start+1) || dfs(x, y-1, start+1)
		used[x][y] = false
		return found
	}
	for i := range board {
		for j := range board[i] {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
