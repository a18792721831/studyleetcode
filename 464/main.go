package main

import "fmt"

func main() {
	maxChoosableInteger := 10
	desiredTotal := 40
	fmt.Println(canIWin(maxChoosableInteger, desiredTotal))
}

// 输入：maxChoosableInteger = 10, desiredTotal = 11
//输出：false
//解释：
//无论第一个玩家选择哪个整数，他都会失败。
//第一个玩家可以选择从 1 到 10 的整数。
//如果第一个玩家选择 1，那么第二个玩家只能选择从 2 到 10 的整数。
//第二个玩家可以通过选择整数 10（那么累积和为 11 >= desiredTotal），从而取得胜利.
//同样地，第一个玩家选择任意其他整数，第二个玩家都会赢。

// A 选一个数 a => a >= target ，a win.  B 选一个数 b，a+b >= target ,b win
// 选过的数不能在选
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	// 这种其实可以直接判定
	// 1 ... 10 = 1+2+3+4+5+6+7+8+9+10 = ((1+10)*10)/2=
	// 1 ... maxChoosableInteger
	if ((1+maxChoosableInteger)*maxChoosableInteger)/2 < desiredTotal {
		return false
	}
	// 特殊值
	if desiredTotal == 0 {
		return true
	}
	// 首先，这是存在性问题，只要找到一条路径，那么就能结束，所以需要短路
	var backtrack func(path []int, sum int) bool
	// 其次，因为选过的数字不能在选，所以需要缓存已经用过的数字
	// 因为 maxChoosableInteger 最大是20，防止溢出，直接设置缓存大小为 maxChoosableInteger + 1
	used := make([]bool, maxChoosableInteger+1)
	// 这里有个核心判断：如果当前值 + 可选列表的最大值，大于 target，那么当前选择玩家赢。
	// 如果当前值 + 可选列表最大值，小于最大值，此时对于当前选择玩家来说，最优选择是最小值？
	backtrack = func(path []int, sum int) bool {
		// 结束条件: sum >= target 或 全部数字用完
		if sum >= desiredTotal {
			// 判断path 元素数量，如果是奇数，那么A 赢，如果是偶数，那么B赢
			if len(path)%2 == 1 {
				return true
			}
			return false
		}
		// 先选择最大的
		for i := maxChoosableInteger; i >= 1; i-- {
			// i已经用过，不能在用
			if used[i] {
				continue
			}
			if sum+i >= desiredTotal {
				// 选择
				path = append(path, i)
				used[i] = true
				// 短路结束
				return backtrack(path, sum+i)
			} else {
				// 跳出，选择最小
				break
			}
		}
		// 选择最小
		for i := 1; i <= maxChoosableInteger; i++ {
			if used[i] {
				continue
			}
			// 选择
			path = append(path, i)
			// 占用
			used[i] = true
			// 递归
			return backtrack(path, sum+i)
		}
		return false
	}
	return backtrack([]int{}, 0)
}
