package main

/*
https://leetcode-cn.com/problems/island-perimeter/
给定一个 row x col 的二维网格地图 grid ，其中：grid[i][j] = 1 表示陆地， grid[i][j] = 0 表示水域。

网格中的格子 水平和垂直 方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。

岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。



示例 1：

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/island-perimeter
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func islandPerimeter(grid [][]int) int {
	l1 := len(grid)
	if l1 == 0 {
		return 0
	}

	ret := 0
	for i1, area := range grid {
		for i2, n := range area {
			if n != 1 {
				continue
			}

			ret += 4
			// 左右重合
			if i2 > 0 {
				if grid[i1][i2-1] == 1 {
					ret -= 2
				}
			}

			if i1 > 0 {
				if grid[i1-1][i2] == 1 { // 上下重合
					ret -= 2
				}
			}
		}
	}

	return ret
}
