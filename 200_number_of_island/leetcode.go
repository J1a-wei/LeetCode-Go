package leetcode

/*

给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

var DIRECTIONS = [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

// 这题注意LeetCode输入参数是一个byte '1' 而并非 1
func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	graph := grid
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		colSlice := make([]bool, cols)
		visited[i] = colSlice
	}

	var inArea func(int, int) bool
	inArea = func(x int, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	var dfs func(int, int)
	dfs = func(x int, y int) {
		visited[x][y] = true
		for k := 0; k < 4; k++ {
			newX := x + DIRECTIONS[k][0]
			newY := y + DIRECTIONS[k][1]
			if inArea(newX, newY) && graph[newX][newY] == '1' && !visited[newX][newY] {
				dfs(newX, newY)
			}
		}
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 岛屿是其中的一个点，并且没有被访问过，就用深度优先遍历
			if !visited[i][j] && graph[i][j] == '1' {
				dfs(i, j)
				count++
			}
		}
	}
	return count
}

// BFS
func numIslands2(grid [][]byte) int {
	rows := len(grid)
	count := 0
	if rows == 0 {
		return count
	}

	cols := len(grid[0])

	visited := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	var inArea func(int, int) bool
	inArea = func(x int, y int) bool {
		return x >= 0 && x < rows && y >= 0 && y < cols
	}

	var bfs func(int, int)
	bfs = func(x int, y int) {
		q := []int{x*cols + y} // 如果用一个节点去存两个节点？ 特殊处理一下

		for len(q) != 0 {
			cur := q[0]
			curX := cur / cols
			curY := cur % cols
			q = q[1:]

			for k := 0; k < 4; k++ {
				nextX := curX + DIRECTIONS[k][0]
				nextY := curY + DIRECTIONS[k][1]
				if inArea(nextX, nextY) && !visited[nextX][nextY] && grid[nextX][nextY] == '1' {
					visited[nextX][nextY] = true
					q = append(q, nextX*cols+nextY)
				}

			}

		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				count++
				bfs(i, j)
			}
		}
	}
	return count
}
